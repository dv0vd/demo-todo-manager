package routes

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/http/middleware"
	"demo-todo-manager/pkg/logger"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterPrivateRoutes(router *chi.Mux, authController contracts.AuthController, noteController contracts.NoteController) {
	logger.Log.Info("Starting registering private routes")

	router.Group(func(private chi.Router) {
		private.Use(middleware.AuthMiddleware)

		private.Post("/auth/refresh", http.HandlerFunc(authController.RefreshToken))

		private.Route("/notes", func(notes chi.Router) {
			notes.Get("/", http.HandlerFunc(noteController.GetAll))
		})
	})

	logger.Log.Info("Private routes have been registered successfully")
}
