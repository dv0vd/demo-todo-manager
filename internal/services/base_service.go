package services

import (
	"demo-todo-manager/internal/contracts"
	"os"
	"strconv"
)

func InitAuthService() contracts.AuthService {
	jwtTtl, _ := strconv.ParseUint(os.Getenv("JWT_TTL"), 10, 0)
	jwtRefreshTtl, _ := strconv.ParseUint(os.Getenv("JWT_REFRESH_TTL"), 10, 0)

	return newAuthService(os.Getenv("JWT_SECRET"), jwtTtl, jwtRefreshTtl)
}

func InitEnvService() contracts.EnvService {
	return &envService{}
}

func InitServices() (contracts.EnvService, contracts.UserService, contracts.DBService, contracts.AuthService, contracts.NoteService) {
	return InitEnvService(),
		InitUserService(true),
		newDBService(
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		),
		InitAuthService(),
		newNoteService(true)
}

func InitUserService(repository bool) contracts.UserService {
	return newUserService(repository)
}
