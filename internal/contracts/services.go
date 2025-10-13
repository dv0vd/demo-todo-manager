package contracts

import (
	"context"
	"demo-todo-manager/internal/dto"

	"github.com/golang-jwt/jwt/v5"
)

type UserIdContextKey string

type AuthService interface {
	ExtractEncodedTokenFromHeader(string) string
	GetToken(string) (*jwt.Token, error)
	GetUserIdFromContext(context.Context) uint64
	GetUserIdKey() UserIdContextKey
	IssueToken(uint64) (string, error)
}

type DBService interface {
	Migrate()
	CloseConnections(UserService, NoteService)
}

type EnvService interface {
	Validate() bool
}

type NoteService interface {
	CloseDBConnection()
	Get(id uint64, userId uint64) (dto.NoteDTO, bool)
	GetByUserId(userId uint64) ([]dto.NoteDTO, bool)
	Delete(id uint64, userId uint64) bool
}

type UserService interface {
	CloseDBConnection()
	GetByEmail(string) (dto.UserDTO, bool)
	GetById(uint64) (dto.UserDTO, bool)
	HashPassword(userDTO dto.UserDTO) (string, error)
	Store(dto.UserDTO) (dto.UserDTO, error)
	ValidatePassword(password, hashedPassword string) bool
}
