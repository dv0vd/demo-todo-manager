package routes

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/pkg/logger"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func registerPublicRoutes(router *chi.Mux, userController contracts.UserController) {
	logger.Log.Info("Starting registering public routes")

	router.Post("/login", http.HandlerFunc(userController.Login))
	router.Post("/signup", http.HandlerFunc(userController.Signup))

	logger.Log.Info("Public routes have been registered successfully")
}
