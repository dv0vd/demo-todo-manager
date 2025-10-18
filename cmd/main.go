package main

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/http/controllers"
	"demo-todo-manager/internal/http/middleware"
	"demo-todo-manager/internal/http/routes"
	"demo-todo-manager/internal/services"
	"demo-todo-manager/pkg/logger"
	"net/http"

	"github.com/go-chi/chi/v5"

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

	userController := controllers.NewUserController(userService, authService)
	authController := controllers.NewAuthController(authService)
	noteController := controllers.NewNoteController(authService, userService, noteService)

	http.ListenAndServe(":8080", initRouter(userController, noteController, authController))

	// todo - graceful shutdown
	// dbService.CloseConnections(userService)
}

func initRouter(userController contracts.UserController, noteController contracts.NoteController, authController contracts.AuthController) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.ContentTypeMiddleware)
	router.Use(middleware.LocaleMiddleware)
	routes.RegisterPublicRoutes(router, userController)
	routes.RegisterPrivateRoutes(router, authController, noteController)
	router.Mount("/api", router)

	return router
}
