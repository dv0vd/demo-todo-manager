package main

import (
	"demo-todo-manager/internal/http/routes"
	"demo-todo-manager/internal/services"
	"demo-todo-manager/pkg/logger"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	logger.Log.Infof("The server is starting...")

	envservice := services.NewEnvService()
	if !envservice.Validate() {
		logger.Log.Fatal("Error starting server: env validation failed")
	}

	m, err := migrate.New(
		"file://migrations/postgres",
		fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", os.Getenv("TODO_MANAGER_DB_USER"), os.Getenv("TODO_MANAGER_DB_PASSWORD"), os.Getenv("TODO_MANAGER_DB_HOST"), os.Getenv("TODO_MANAGER_DB_PORT"), os.Getenv("TODO_MANAGER_DB_NAME")),
	)
	if err != nil {
		logger.Log.Fatal(fmt.Sprintf("Error starting server: connection to DB failed. Error: %v", err.Error()))
	}
	if err := m.Up(); err != nil {
		logger.Log.Fatal(fmt.Sprintf("Error starting server: migrations failed. Error: %v", err.Error()))
	}

	mux := http.NewServeMux()
	routes := routes.RegisterPublicRoutes(mux)
	http.ListenAndServe(":8080", routes)
}
