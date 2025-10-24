package routes

import (
	"demo-todo-manager/docs"
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/http/middleware"
	"os"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/chi/v5"
)

func InitRouter(
	userController contracts.UserController,
	noteController contracts.NoteController,
	authController contracts.AuthController,
) *chi.Mux {
	router := chi.NewRouter()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.Handle("/swagger/*", httpSwagger.WrapHandler)

	api := chi.NewRouter()

	api.Use(middleware.ContentTypeMiddleware)
	api.Use(middleware.LocaleMiddleware)
	registerPublicRoutes(api, userController)
	registerPrivateRoutes(api, authController, noteController)
	router.Mount("/api", api)

	return router
}
