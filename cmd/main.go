package main

import (
	"context"
	"demo-todo-manager/internal/http/controllers"
	"demo-todo-manager/internal/http/routes"
	"demo-todo-manager/internal/services"
	"demo-todo-manager/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "demo-todo-manager/docs"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// @title Demo Todo Manager API
// @version 1.0
// @description Demo REST API for todo management app.
// @host localhost:8080
// @BasePath /api

// @contact.name Viacheslav Davydov
// @contact.url https://dv0vd.dev
// @contact.email viacheslav.davydov@dv0vd.xyz

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
func main() {
	logger.Log.Infof("The server is starting...")

	envService, userService, dbService, authService, noteService := services.InitServices()

	if !envService.Validate() {
		logger.Log.Fatal("Error starting server: env validation failed")
	}

	dbService.Migrate()

	userController, authController, noteController := controllers.InitControllers(userService, authService, noteService)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes.InitRouter(userController, noteController, authController),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.WithField("error", err).Fatalf("Error starting server: %v", err.Error())
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	logger.Log.Infof("Shutting down the server...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Log.WithField("error", err).Errorf("Server Shutdown Failed: %v", err.Error())
	}

	dbService.CloseConnections(userService, noteService)
	logger.Log.Infof("Server exited successfully")
}
