package contracts

import "demo-todo-manager/internal/dto"

type NoteRepository interface {
	CloseDBConnection()
	Create(dto.NoteDTO, int64) (dto.NoteDTO, error)
	Get(uint64, uint64) (dto.NoteDTO, bool)
	GetByUserId(uint64) ([]dto.NoteDTO, bool)
	Delete(uint64, uint64) bool
	Update(dto.NoteDTO, uint64) bool
}

type UserRepository interface {
	CloseDBConnection()
	GetByEmail(string) (dto.UserDTO, bool)
	GetById(uint64) (dto.UserDTO, bool)
	Store(dto.UserDTO) (dto.UserDTO, error)
}
