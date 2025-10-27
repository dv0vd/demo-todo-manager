package repositories

import (
	"database/sql"
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/dto"
	"demo-todo-manager/pkg/logger"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type noteRepository struct {
	client     *sql.DB
	table      string
	filterable []string
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
		client:     db,
		table:      "notes",
		filterable: []string{"title", "description", "done"},
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
			"INSERT INTO %v(title, description, done, user_id) VALUES($1, $2, $3, $4) RETURNING id, created_at, updated_at",
			r.table,
		),
		noteDTO.Title,
		noteDTO.Description,
		noteDTO.Done,
		userId,
	).Scan(&noteDTO.ID, &noteDTO.CreatedAt, &noteDTO.UpdatedAt); err != nil {
		logger.Log.WithFields(logrus.Fields{"noteDTO": noteDTO, "userId": userId}).Errorf("Error during note creation: %v", err.Error())

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

		logger.Log.WithFields(logrus.Fields{"noteId": id, "userId": userId}).Errorf("Failed getting note by id '%v'. Error: %v", id, err.Error())

		return dto.NoteDTO{}, false
	}

	return noteDTO, true
}

func (r *noteRepository) GetByUserId(userId uint64, filters map[string]interface{}) ([]dto.NoteDTO, bool) {
	notes := []dto.NoteDTO{}

	whereClauses := []string{"user_id=$1"}
	whereArgs := []interface{}{userId}

	i := 2
	for key, where := range filters {
		for _, allowed := range r.filterable {
			if key == allowed {
				whereClauses = append(whereClauses, fmt.Sprintf("%v=$%v", key, i))
				whereArgs = append(whereArgs, where)

				continue
			}
		}
	}

	query := fmt.Sprintf("SELECT id, title, description, created_at, updated_at FROM %v WHERE %v", r.table, strings.Join(whereClauses, " AND "))
	rows, err := r.client.Query(query, whereArgs...)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{"userId": userId, "filters": filters, "args": whereArgs, "query": query}).Errorf("Failed getting notes by user id '%v'. Error: %v", userId, err.Error())

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
	result, err := r.client.Exec(
		fmt.Sprintf(
			"UPDATE %v SET title=$1, description=$2, done=$3, updated_at=NOW() WHERE id=$4 AND user_id=$5",
			r.table,
		),
		noteDTO.Title,
		noteDTO.Description,
		noteDTO.Done,
		noteDTO.ID,
		userId,
	)

	if err != nil {
		logger.Log.WithFields(logrus.Fields{"noteDTO": noteDTO, "userId": userId}).Errorf("Failed updating note #%v. Error: %v", noteDTO.ID, err.Error())

		return false
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Log.WithFields(logrus.Fields{"noteDTO": noteDTO, "userId": userId}).Errorf("Failed obtaining affected rows amount for note #%v. Error: %v", noteDTO.ID, err.Error())

		return false
	}

	if rowsAffected == 0 {
		return false
	}

	return true
}
