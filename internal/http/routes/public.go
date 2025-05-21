package routes

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/http/middleware"
	"demo-todo-manager/pkg/logger"
	"net/http"
)

func RegisterPublicRoutes(mux *http.ServeMux, userController contracts.UserController) http.Handler {
	logger.Log.Info("Starting registering public routes")

	mux.HandleFunc("/api/signup", userController.Signup)

	contentTypeMux :=
		middleware.ContentTypeMiddleware(mux)

	logger.Log.Info("Public routes have been registered successfully")

	return contentTypeMux
}
