package repositories

import (
	"database/sql"
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/dto"
	"demo-todo-manager/pkg/logger"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type noteRepository struct {
	client *sql.DB
	table  string
}

func NewNoteRepositoryPostgres() contracts.NoteRepository {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		logger.Log.Fatal(fmt.Sprintf("Error creating connection to note repository. Error: %v", err.Error()))
	}

	err = db.Ping()
	if err != nil {
		logger.Log.Fatal(fmt.Sprintf("Error testing connection to note repository. Error: %v", err.Error()))
	}

	return &noteRepository{
		client: db,
		table:  "notes",
	}
}

func (r *noteRepository) CloseDBConnection() {
	if r.client != nil {
		r.client.Close()
	}
}

func (r *noteRepository) Create(noteDTO dto.NoteDTO, userId uint64) (dto.NoteDTO, error) {
	if err := r.client.QueryRow(
		fmt.Sprintf(
			"INSERT INTO %v(title, description, user_id) VALUES($1, $2, $3) RETURNING id, created_at, updated_at",
			r.table,
		),
		noteDTO.Title,
		noteDTO.Description,
		userId,
	).Scan(&noteDTO.ID, &noteDTO.CreatedAt, &noteDTO.UpdatedAt); err != nil {
		logger.Log.WithFields(logrus.Fields{"noteDTO": noteDTO}).Errorf("Error during note creation: %v", err.Error())

		return noteDTO, err
	}

	return noteDTO, nil
}

func (r *noteRepository) Get(id uint64, userId uint64) (dto.NoteDTO, bool) {
	var noteDTO dto.NoteDTO

	if err := r.client.QueryRow(
		fmt.Sprintf(
			"SELECT id, title, description, created_at, updated_at FROM %v WHERE id=$1 AND user_id=$2",
			r.table),
		id,
		userId,
	).Scan(
		&noteDTO.ID, &noteDTO.Title, &noteDTO.Description, &noteDTO.CreatedAt, &noteDTO.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return noteDTO, true
		}

		logger.Log.WithFields(logrus.Fields{"noteId": id}).Errorf("Failed getting note by id '%v'. Error: %v", id, err.Error())

		return dto.NoteDTO{}, false
	}

	return noteDTO, true
}

func (r *noteRepository) GetByUserId(userId uint64) ([]dto.NoteDTO, bool) {
	notes := []dto.NoteDTO{}

	rows, err := r.client.Query(fmt.Sprintf("SELECT id, title, description, created_at, updated_at FROM %v WHERE user_id=$1", r.table), userId)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{"userId": userId}).Errorf("Failed getting notes by user id '%v'. Error: %v", userId, err.Error())

		return []dto.NoteDTO{}, false
	}
	defer rows.Close()

	for rows.Next() {
		var noteDTO dto.NoteDTO

		err = rows.Scan(&noteDTO.ID, &noteDTO.Title, &noteDTO.Description, &noteDTO.CreatedAt, &noteDTO.UpdatedAt)
		if err != nil {
			logger.Log.WithFields(logrus.Fields{"userId": userId}).Errorf("Failed getting notes by user id '%v'. Error: %v", userId, err.Error())

			return []dto.NoteDTO{}, false
		}

		notes = append(notes, noteDTO)
	}

	if err = rows.Err(); err != nil {
		logger.Log.WithFields(logrus.Fields{"userId": userId}).Errorf("Failed getting notes by user id '%v'. Error: %v", userId, err.Error())

		return []dto.NoteDTO{}, false
	}

	return notes, true
}

func (r *noteRepository) Delete(id uint64, userId uint64) bool {
	_, err := r.client.Exec(fmt.Sprintf("DELETE FROM %v WHERE id=$1 AND user_id=$2", r.table), id, userId)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{"noteId": id}).Errorf("Failed deleting note by id '%v'. Error: %v", id, err.Error())

		return false
	}

	return true
}

func (r *noteRepository) Update(noteDTO dto.NoteDTO, userId uint64) bool {
	if _, err := r.client.Exec(
		fmt.Sprintf(
			"UPDATE %v SET title=$1, description=$2, updated_at=NOW() WHERE id=$3 AND user_id=$4",
			r.table,
		),
		noteDTO.Title,
		noteDTO.Description,
		noteDTO.ID,
		userId,
	); err != nil {
		logger.Log.WithFields(logrus.Fields{"noteDTO": noteDTO}).Errorf("Failed updating note #%v. Error: %v", noteDTO.ID, err.Error())

		return false
	}

	return true
}
