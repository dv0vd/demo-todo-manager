package services

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/pkg/logger"
	"os"
	"strconv"
)

type envService struct{}

func NewEnvService() contracts.EnvService {
	return &envService{}
}

func (s *envService) Validate() bool {
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

	value = os.Getenv("JWT_TTL")
	if value == "" {
		logger.Log.Error("JWT_TTL is not set!")

		return false
	}
	_, err := strconv.ParseUint(value, 10, 0)
	if err != nil {
		logger.Log.Error("JWT_TTL has incorrect format!")

		return false
	}

	value = os.Getenv("JWT_REFRESH_TTL")
	if value == "" {
		logger.Log.Error("JWT_REFRESH_TTL is not set!")

		return false
	}
	_, err = strconv.ParseUint(value, 10, 0)
	if err != nil {
		logger.Log.Error("JWT_REFRESH_TTL has incorrect format!")

		return false
	}

	value = os.Getenv("JWT_SECRET")
	if value == "" {
		logger.Log.Error("JWT_SECRET is not set!")

		return false
	}

	logger.Log.Info("Env validation finished successfully")

	return true
}
