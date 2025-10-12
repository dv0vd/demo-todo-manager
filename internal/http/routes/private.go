package routes

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/http/middleware"
	"demo-todo-manager/pkg/logger"
	"net/http"
)

func RegisterPrivateRoutes(mux *http.ServeMux, authController contracts.AuthController, noteController contracts.NoteController) {
	logger.Log.Info("Starting registering private routes")

	mux.Handle(
		"/api/auth/refresh",
		middleware.ContentTypeMiddleware(
			middleware.AuthMiddleware(
				http.HandlerFunc(authController.RefreshToken),
				authController.GetAuthService(),
				noteController.GetNoteService(),
			),
		),
	)

	mux.Handle(
		"/api/notes",
		middleware.ContentTypeMiddleware(
			middleware.AuthMiddleware(
				http.HandlerFunc(noteController.GetAll),
				authController.GetAuthService(),
				noteController.GetNoteService(),
			),
		),
	)

	logger.Log.Info("Private routes have been registered successfully")
}
