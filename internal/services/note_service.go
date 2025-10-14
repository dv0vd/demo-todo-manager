package services

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/dto"
	repositories "demo-todo-manager/internal/repositories/postgres"
)

type noteService struct {
	repository contracts.NoteRepository
}

func NewNoteService(repository bool) contracts.NoteService {
	if repository {
		return &noteService{
			repository: repositories.NewNoteRepositoryPostgres(),
		}
	}

	return &noteService{}
}

func (s *noteService) CloseDBConnection() {
	s.repository.CloseDBConnection()
}

func (s *noteService) Get(id uint64, userId uint64) (dto.NoteDTO, bool) {
	return s.repository.Get(id, userId)
}

func (s *noteService) GetByUserId(userId uint64) ([]dto.NoteDTO, bool) {
	return s.repository.GetByUserId(userId)
}

func (s *noteService) Delete(id uint64, userId uint64) bool {
	return s.repository.Delete(id, userId)
}

func (s *noteService) Update(noteDTO dto.NoteDTO, userId uint64) bool {
	return s.repository.Update(noteDTO, userId)
}
