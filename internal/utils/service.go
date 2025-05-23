package utils

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/services"
	"os"
	"strconv"
)

func ServiceInitServices() (contracts.EnvService, contracts.UserService, contracts.DBService, contracts.AuthService) {
	jwtTtl, _ := strconv.ParseUint(os.Getenv("JWT_TTL"), 10, 0)
	jwtRefreshTtl, _ := strconv.ParseUint(os.Getenv("JWT_REFRESH_TTL"), 10, 0)

	return services.NewEnvService(),
		services.NewUserService(true),
		services.NewDBService(
			os.Getenv("TODO_MANAGER_DB_USER"),
			os.Getenv("TODO_MANAGER_DB_PASSWORD"),
			os.Getenv("TODO_MANAGER_DB_HOST"),
			os.Getenv("TODO_MANAGER_DB_PORT"),
			os.Getenv("TODO_MANAGER_DB_NAME"),
		),
		services.NewAuthService(os.Getenv("JWT_SECRET"), jwtTtl, jwtRefreshTtl)
}
