package contracts

import "demo-todo-manager/internal/dto"

type UserRepository interface {
	CloseDBConnection()
	GetByEmail(string) (dto.UserDTO, bool)
}
