package services

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/pkg/logger"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type dbService struct{}

func NewDBService() contracts.DBService {
	// todo - pass env
	return &dbService{}
}

func (s *dbService) CloseConnections(userService contracts.UserService) {
	userService.CloseDBConnection()
}

func (s *dbService) Migrate() {
	m, err := migrate.New(
		"file:///app/migrations/postgres",
		fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", os.Getenv("TODO_MANAGER_DB_USER"), os.Getenv("TODO_MANAGER_DB_PASSWORD"), os.Getenv("TODO_MANAGER_DB_HOST"), os.Getenv("TODO_MANAGER_DB_PORT"), os.Getenv("TODO_MANAGER_DB_NAME")),
	)
	if err != nil {
		logger.Log.Fatalf("Failed to read DB migrations. Error: %v", err.Error())
	}

	logger.Log.Info("Starting DB migrations")
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			logger.Log.Info("All migrations have already been applied")
		} else {
			logger.Log.Fatalf("DB Migrations failed. Error: %v", err.Error())
		}
	}
	logger.Log.Info("DB migrations completed successfully")
}
