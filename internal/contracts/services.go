package contracts

import (
	"demo-todo-manager/internal/dto"
)

type AuthService interface {
	IssueToken(uint64) (string, error)
}

type DBService interface {
	Migrate()
	CloseConnections(UserService)
}

type EnvService interface {
	Validate() bool
}

type UserService interface {
	CloseDBConnection()
	GetByEmail(string) (dto.UserDTO, bool)
	HashPassword(userDTO dto.UserDTO) (string, error)
	Store(dto.UserDTO) (dto.UserDTO, error)
	ValidatePassword(password, hashedPassword string) bool
}
