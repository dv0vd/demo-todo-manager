package services

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/pkg/logger"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type dbService struct {
	user     string
	password string
	host     string
	port     string
	database string
}

func NewDBService(
	user,
	password,
	host,
	port,
	database string,
) contracts.DBService {
	return &dbService{
		user:     user,
		password: password,
		host:     host,
		port:     port,
		database: database,
	}
}

func (s *dbService) CloseConnections(
	userService contracts.UserService,
	noteService contracts.NoteService,
) {
	userService.CloseDBConnection()
	noteService.CloseDBConnection()
}

func (s *dbService) Migrate() {
	m, err := migrate.New(
		"file:///app/migrations/postgres",
		fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", s.user, s.password, s.host, s.port, s.database),
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
