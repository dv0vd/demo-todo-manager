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
