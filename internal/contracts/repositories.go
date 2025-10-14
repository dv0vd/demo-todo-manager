package contracts

import "demo-todo-manager/internal/dto"

type NoteRepository interface {
	CloseDBConnection()
	Get(id uint64, userId uint64) (dto.NoteDTO, bool)
	GetByUserId(userId uint64) ([]dto.NoteDTO, bool)
	Delete(id uint64, userId uint64) bool
	Update(noteDTO dto.NoteDTO, userId uint64) bool
}

type UserRepository interface {
	CloseDBConnection()
	GetByEmail(string) (dto.UserDTO, bool)
	GetById(uint64) (dto.UserDTO, bool)
	Store(dto.UserDTO) (dto.UserDTO, error)
}
