package main

import (
	"demo-todo-manager/internal/http/routes"
	"demo-todo-manager/internal/utils"
	"demo-todo-manager/pkg/logger"
	"net/http"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	logger.Log.Infof("The server is starting...")

	envService, userService, dbService, authService := utils.ServiceInitServices()

	if !envService.Validate() {
		logger.Log.Fatal("Error starting server: env validation failed")
	}

	dbService.Migrate()

	userController := utils.ControllerInitControllers(userService, authService)

	mux := http.NewServeMux()
	routes := routes.RegisterPublicRoutes(mux, userController)
	http.ListenAndServe(":8080", routes)

	// todo - graceful shutdown
	// dbService.CloseConnections(userService)
}
