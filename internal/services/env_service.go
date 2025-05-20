package services

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/pkg/logger"
	"os"
)

type EnvService struct{}

func NewEnvService() contracts.EnvService {
	return &EnvService{}
}

func (e *EnvService) Validate() bool {
	value := os.Getenv("TODO_MANAGER_DB_HOST")
	if value == "" {
		logger.Log.Error("TODO_MANAGER_DB_HOST is not set!")

		return false
	}

	value = os.Getenv("TODO_MANAGER_DB_PORT")
	if value == "" {
		logger.Log.Error("TODO_MANAGER_DB_PORT is not set!")

		return false
	}

	value = os.Getenv("TODO_MANAGER_DB_NAME")
	if value == "" {
		logger.Log.Error("TODO_MANAGER_DB_NAME is not set!")

		return false
	}

	value = os.Getenv("TODO_MANAGER_DB_USER")
	if value == "" {
		logger.Log.Error("TODO_MANAGER_DB_USER is not set!")

		return false
	}

	logger.Log.Info("Env validation finished successfully")

	return true
}
