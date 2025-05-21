package contracts

import (
	"demo-todo-manager/internal/dto"
)

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
}
