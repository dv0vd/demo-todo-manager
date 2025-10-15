package main

import (
	"demo-todo-manager/internal/http/controllers"
	"demo-todo-manager/internal/http/routes"
	"demo-todo-manager/internal/utils"
	"demo-todo-manager/pkg/logger"
	"net/http"

	"github.com/go-chi/chi/v5"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	logger.Log.Infof("The server is starting...")

	envService, userService, dbService, authService, noteService := utils.ServiceInitServices()

	if !envService.Validate() {
		logger.Log.Fatal("Error starting server: env validation failed")
	}

	dbService.Migrate()

	userController := controllers.NewUserController(userService, authService)
	authController := controllers.NewAuthController(authService)
	noteController := controllers.NewNoteController(authService, userService, noteService)

	router := chi.NewRouter()
	api := chi.NewRouter()
	routes.RegisterPublicRoutes(api, userController)
	routes.RegisterPrivateRoutes(api, authController, noteController)
	router.Mount("/api", api)

	http.ListenAndServe(":8080", router)

	// todo - graceful shutdown
	// dbService.CloseConnections(userService)
}
