package routes

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/http/middleware"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/chi/v5"
)

func InitRouter(
	userController contracts.UserController,
	noteController contracts.NoteController,
	authController contracts.AuthController,
) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	api := chi.NewRouter()

	api.Use(middleware.ContentTypeMiddleware)
	api.Use(middleware.LocaleMiddleware)
	registerPublicRoutes(api, userController)
	registerPrivateRoutes(api, authController, noteController)
	api.Mount("/api", api)
	router.Mount("/api", api)

	return router
}
