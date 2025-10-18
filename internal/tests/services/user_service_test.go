package tests

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/dto"
	"demo-todo-manager/internal/services"
	testutils "demo-todo-manager/internal/tests"
	"testing"

	"github.com/go-faker/faker/v4"
)

type passwords struct {
	password       string
	hashedPassword string
}

func TestUserServiceValidatePassword(t *testing.T) {
	userService := services.NewUserService(false)

	tests := []struct {
		name      string
		passwords passwords
		expected  bool
	}{
		{
			name:      "Password is correct",
			passwords: generatePasswordsPair(userService, true),
			expected:  true,
		},
		{
			name:      "Password is incorrect",
			passwords: generatePasswordsPair(userService, false),
			expected:  false,
		},
	}

	for _, test := range tests {
		result := userService.ValidatePassword(test.passwords.password, test.passwords.hashedPassword)
		testutils.CheckResult(t, test.name, test.expected, result)
	}
}

func generatePasswordsPair(userService contracts.UserService, correct bool) passwords {
	password := faker.Word()
	var userDTO dto.UserDTO

	if correct {
		userDTO.Password = password
	} else {
		userDTO.Password = faker.Word()
	}

	hashedPassword, _ := userService.HashPassword(userDTO)

	passwords := passwords{
		password:       password,
		hashedPassword: hashedPassword,
	}

	return passwords
}
