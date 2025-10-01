package routes

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/http/middleware"
	"demo-todo-manager/pkg/logger"
	"net/http"
)

func RegisterPublicRoutes(mux *http.ServeMux, userController contracts.UserController) {
	logger.Log.Info("Starting registering public routes")

	mux.Handle("/api/signup", middleware.ContentTypeMiddleware(http.HandlerFunc(userController.Signup)))
	mux.Handle("/api/login", middleware.ContentTypeMiddleware(http.HandlerFunc(userController.Login)))

	logger.Log.Info("Public routes have been registered successfully")
}
