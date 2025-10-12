package contracts

import "demo-todo-manager/internal/dto"

type NoteRepository interface {
	CloseDBConnection()
	GetByUserId(userId uint64) ([]dto.NoteDTO, bool)
}

type UserRepository interface {
	CloseDBConnection()
	GetByEmail(string) (dto.UserDTO, bool)
	GetById(uint64) (dto.UserDTO, bool)
	Store(dto.UserDTO) (dto.UserDTO, error)
}
