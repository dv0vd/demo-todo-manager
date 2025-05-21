package services

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/dto"
	repositories "demo-todo-manager/internal/repositories/postgres"
)

type userService struct {
	repository contracts.UserRepository
}

func NewUserService() contracts.UserService {
	return &userService{
		repository: repositories.NewUserRepositoryPostgres(),
	}
}

func (s *userService) CloseDBConnection() {
	s.repository.CloseDBConnection()
}

func (s *userService) GetByEmail(email string) (dto.UserDTO, bool) {
	return s.repository.GetByEmail(email)
}
