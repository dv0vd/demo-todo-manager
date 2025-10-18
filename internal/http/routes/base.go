package routes

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/http/middleware"

	"github.com/go-chi/chi/v5"
)

func InitRouter(
	userController contracts.UserController,
	noteController contracts.NoteController,
	authController contracts.AuthController,
) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.ContentTypeMiddleware)
	router.Use(middleware.LocaleMiddleware)
	registerPublicRoutes(router, userController)
	registerPrivateRoutes(router, authController, noteController)
	router.Mount("/api", router)

	return router
}
