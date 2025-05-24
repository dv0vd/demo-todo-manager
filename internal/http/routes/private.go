package routes

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/http/middleware"
	"demo-todo-manager/pkg/logger"
	"net/http"
)

func RegisterPrivateRoutes(mux *http.ServeMux, userController contracts.UserController) http.Handler {
	logger.Log.Info("Starting registering private routes")

	// mux.HandleFunc("/api/signup", userController.Signup)
	// mux.HandleFunc("/api/login", userController.Login)

	contentTypeMux :=
		middleware.ContentTypeMiddleware(mux)
	authMux := middleware.AuthMiddleware(contentTypeMux, userController.GetAuthService())

	logger.Log.Info("Private routes have been registered successfully")

	return authMux
}
