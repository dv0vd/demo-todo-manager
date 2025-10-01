package contracts

import (
	"demo-todo-manager/internal/dto"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	ExtractEncodedTokenFromHeader(string) string
	GetToken(string) (*jwt.Token, error)
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
