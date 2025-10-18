package main

import (
	"demo-todo-manager/internal/http/controllers"
	"demo-todo-manager/internal/http/routes"
	"demo-todo-manager/internal/services"
	"demo-todo-manager/pkg/logger"
	"net/http"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	logger.Log.Infof("The server is starting...")

	envService, userService, dbService, authService, noteService := services.InitServices()

	if !envService.Validate() {
		logger.Log.Fatal("Error starting server: env validation failed")
	}

	dbService.Migrate()

	userController, authController, noteController := controllers.InitControllers(userService, authService, noteService)

	http.ListenAndServe(":8080", routes.InitRouter(userController, noteController, authController))

	// todo - graceful shutdown
	// dbService.CloseConnections(userService)
}
