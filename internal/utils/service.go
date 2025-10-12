package utils

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/services"
	"os"
	"strconv"
)

func ServiceInitServices() (contracts.EnvService, contracts.UserService, contracts.DBService, contracts.AuthService, contracts.NoteService) {
	return services.NewEnvService(),
		services.NewUserService(true),
		services.NewDBService(
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		),
		ServicesInitAuthService(),
		services.NewNoteService(true)
}

func ServicesInitAuthService() contracts.AuthService {
	jwtTtl, _ := strconv.ParseUint(os.Getenv("JWT_TTL"), 10, 0)
	jwtRefreshTtl, _ := strconv.ParseUint(os.Getenv("JWT_REFRESH_TTL"), 10, 0)

	return services.NewAuthService(os.Getenv("JWT_SECRET"), jwtTtl, jwtRefreshTtl)
}
