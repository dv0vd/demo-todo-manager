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

type userRepository struct {
	client *sql.DB
	table  string
}

func NewUserRepositoryPostgres() contracts.UserRepository {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", os.Getenv("TODO_MANAGER_DB_USER"), os.Getenv("TODO_MANAGER_DB_PASSWORD"), os.Getenv("TODO_MANAGER_DB_HOST"), os.Getenv("TODO_MANAGER_DB_PORT"), os.Getenv("TODO_MANAGER_DB_NAME")))
	if err != nil {
		logger.Log.Fatal(fmt.Sprintf("Error creating connection to user repository. Error: %v", err.Error()))
	}

	err = db.Ping()
	if err != nil {
		logger.Log.Fatal(fmt.Sprintf("Error testing connection to user repository. Error: %v", err.Error()))
	}

	return &userRepository{
		client: db,
		table:  "users",
	}
}

func (r *userRepository) CloseDBConnection() {
	if r.client != nil {
		r.client.Close()
	}
}

func (r *userRepository) GetByEmail(email string) (dto.UserDTO, bool) {
	var userDTO dto.UserDTO

	if err := r.client.QueryRow(fmt.Sprintf("SELECT * FROM %v WHERE email=$1", r.table), email).Scan(&userDTO.ID, &userDTO.Email, &userDTO.Password); err != nil {
		logger.Log.Warningf("Failed getting user by email '%v'. Error: %v", email, err.Error())

		if err == sql.ErrNoRows {
			return userDTO, true
		}

		return userDTO, false
	}

	return userDTO, true
}

func (r *userRepository) Store(userDTO dto.UserDTO) (dto.UserDTO, error) {
	if err := r.client.QueryRow(fmt.Sprintf("INSERT INTO %v (email, password) VALUES ($1, $2) RETURNING id", r.table), userDTO.Email, userDTO.Password).Scan(&userDTO.ID); err != nil {
		logger.Log.WithFields(logrus.Fields{"userDTO": userDTO}).Warningf("Failed inserting user '%v'. Error: %v", userDTO.Email, err.Error())

		return userDTO, err
	}

	return userDTO, nil
}
