package services

import (
	"demo-todo-manager/internal/contracts"
	"os"
	"strconv"
)

func InitAuthService() contracts.AuthService {
	jwtTtl, _ := strconv.ParseUint(os.Getenv("JWT_TTL"), 10, 0)
	jwtRefreshTtl, _ := strconv.ParseUint(os.Getenv("JWT_REFRESH_TTL"), 10, 0)

	return NewAuthService(os.Getenv("JWT_SECRET"), jwtTtl, jwtRefreshTtl)
}

func InitServices() (contracts.EnvService, contracts.UserService, contracts.DBService, contracts.AuthService, contracts.NoteService) {
	return NewEnvService(),
		NewUserService(true),
		NewDBService(
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		),
		InitAuthService(),
		NewNoteService(true)
}
