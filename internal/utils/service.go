package utils

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/services"
)

func ServiceInitServices() (contracts.EnvService, contracts.UserService, contracts.DBService) {
	return services.NewEnvService(), services.NewUserService(), services.NewDBService()
}
