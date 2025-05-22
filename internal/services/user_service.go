package services

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/dto"
	repositories "demo-todo-manager/internal/repositories/postgres"
	"demo-todo-manager/pkg/logger"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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

func (s *userService) Store(userDTO dto.UserDTO) (dto.UserDTO, error) {
	hashedPassword, err := s.hashPassword(userDTO)
	if err != nil {
		return userDTO, err
	}

	userDTO.Password = hashedPassword

	return s.repository.Store(userDTO)
}

func (s *userService) hashPassword(userDTO dto.UserDTO) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)

	if err != nil {
		logger.Log.WithFields(logrus.Fields{"userDTO": userDTO}).Warningf("Error hashing password for user '%v'", userDTO.Email)

		return "", err
	}

	return string(hashedPassword), nil
}
