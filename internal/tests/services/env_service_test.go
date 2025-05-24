package tests

import (
	"demo-todo-manager/internal/services"
	"demo-todo-manager/internal/utils"
	"os"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestEnvServiceValidate(t *testing.T) {
	tests := []struct {
		name     string
		env      map[string]string
		expected bool
	}{
		{
			name: "All env variables are correctly set",
			env: map[string]string{
				"TODO_MANAGER_DB_HOST": "localhost",
				"TODO_MANAGER_DB_PORT": "5432",
				"TODO_MANAGER_DB_NAME": "name",
				"TODO_MANAGER_DB_USER": "user",
				"JWT_TTL":              utils.TestGetRandomInt(1, 1000),
				"JWT_REFRESH_TTL":      utils.TestGetRandomInt(1, 1000),
				"JWT_SECRET":           faker.Word(),
			},
			expected: true,
		},
		{
			name: "TODO_MANAGER_DB_HOST env variable is not set",
			env: map[string]string{
				"TODO_MANAGER_DB_PORT": "5432",
				"TODO_MANAGER_DB_NAME": "name",
				"TODO_MANAGER_DB_USER": "user",
				"JWT_TTL":              utils.TestGetRandomInt(1, 1000),
				"JWT_REFRESH_TTL":      utils.TestGetRandomInt(1, 1000),
				"JWT_SECRET":           faker.Word(),
			},
			expected: false,
		},
		{
			name: "TODO_MANAGER_DB_PORT env variable is not set",
			env: map[string]string{
				"TODO_MANAGER_DB_HOST": "localhost",
				"TODO_MANAGER_DB_NAME": "name",
				"TODO_MANAGER_DB_USER": "user",
				"JWT_TTL":              utils.TestGetRandomInt(1, 1000),
				"JWT_REFRESH_TTL":      utils.TestGetRandomInt(1, 1000),
				"JWT_SECRET":           faker.Word(),
			},
			expected: false,
		},
		{
			name: "TODO_MANAGER_DB_NAME env variable is not set",
			env: map[string]string{
				"TODO_MANAGER_DB_HOST": "localhost",
				"TODO_MANAGER_DB_PORT": "5432",
				"TODO_MANAGER_DB_USER": "user",
				"JWT_TTL":              utils.TestGetRandomInt(1, 1000),
				"JWT_REFRESH_TTL":      utils.TestGetRandomInt(1, 1000),
				"JWT_SECRET":           faker.Word(),
			},
			expected: false,
		},
		{
			name: "TODO_MANAGER_DB_USER env variable is not set",
			env: map[string]string{
				"TODO_MANAGER_DB_HOST": "localhost",
				"TODO_MANAGER_DB_PORT": "5432",
				"TODO_MANAGER_DB_NAME": "name",
				"JWT_TTL":              utils.TestGetRandomInt(1, 1000),
				"JWT_REFRESH_TTL":      utils.TestGetRandomInt(1, 1000),
				"JWT_SECRET":           faker.Word(),
			},
			expected: false,
		},
		{
			name: "JWT_TTL env variable is not set",
			env: map[string]string{
				"TODO_MANAGER_DB_HOST": "localhost",
				"TODO_MANAGER_DB_PORT": "5432",
				"TODO_MANAGER_DB_NAME": "name",
				"TODO_MANAGER_DB_USER": "user",
				"JWT_REFRESH_TTL":      utils.TestGetRandomInt(1, 1000),
				"JWT_SECRET":           faker.Word(),
			},
			expected: false,
		},
		{
			name: "JWT_REFRESH_TTL env variable is not set",
			env: map[string]string{
				"TODO_MANAGER_DB_HOST": "localhost",
				"TODO_MANAGER_DB_PORT": "5432",
				"TODO_MANAGER_DB_NAME": "name",
				"TODO_MANAGER_DB_USER": "user",
				"JWT_TTL":              utils.TestGetRandomInt(1, 1000),
				"JWT_SECRET":           faker.Word(),
			},
			expected: false,
		},
		{
			name: "JWT_SECRET env variable is not set",
			env: map[string]string{
				"TODO_MANAGER_DB_HOST": "localhost",
				"TODO_MANAGER_DB_PORT": "5432",
				"TODO_MANAGER_DB_NAME": "name",
				"TODO_MANAGER_DB_USER": "user",
				"JWT_TTL":              utils.TestGetRandomInt(1, 1000),
				"JWT_REFRESH_TTL":      utils.TestGetRandomInt(1, 1000),
			},
			expected: false,
		},
		{
			name: "JWT_TTL env variable is negative",
			env: map[string]string{
				"TODO_MANAGER_DB_HOST": "localhost",
				"TODO_MANAGER_DB_PORT": "5432",
				"TODO_MANAGER_DB_NAME": "name",
				"TODO_MANAGER_DB_USER": "user",
				"JWT_TTL":              utils.TestGetRandomInt(-1000, -1),
				"JWT_REFRESH_TTL":      utils.TestGetRandomInt(1, 1000),
				"JWT_SECRET":           faker.Word(),
			},
			expected: false,
		},
		{
			name: "JWT_REFRESH_TTL env variable is negative",
			env: map[string]string{
				"TODO_MANAGER_DB_HOST": "localhost",
				"TODO_MANAGER_DB_PORT": "5432",
				"TODO_MANAGER_DB_NAME": "name",
				"TODO_MANAGER_DB_USER": "user",
				"JWT_TTL":              utils.TestGetRandomInt(1, 1000),
				"JWT_REFRESH_TTL":      utils.TestGetRandomInt(-1000, -1),
				"JWT_SECRET":           faker.Word(),
			},
			expected: false,
		},
		{
			name: "JWT_TTL env variable is string",
			env: map[string]string{
				"TODO_MANAGER_DB_HOST": "localhost",
				"TODO_MANAGER_DB_PORT": "5432",
				"TODO_MANAGER_DB_NAME": "name",
				"TODO_MANAGER_DB_USER": "user",
				"JWT_TTL":              faker.Word(),
				"JWT_REFRESH_TTL":      utils.TestGetRandomInt(1, 1000),
				"JWT_SECRET":           faker.Word(),
			},
			expected: false,
		},
		{
			name: "JWT_REFRESH_TTL env variable is string",
			env: map[string]string{
				"TODO_MANAGER_DB_HOST": "localhost",
				"TODO_MANAGER_DB_PORT": "5432",
				"TODO_MANAGER_DB_NAME": "name",
				"TODO_MANAGER_DB_USER": "user",
				"JWT_TTL":              utils.TestGetRandomInt(1, 1000),
				"JWT_REFRESH_TTL":      faker.Word(),
				"JWT_SECRET":           faker.Word(),
			},
			expected: false,
		},
	}

	envservice := services.NewEnvService()

	for _, test := range tests {
		for key, value := range test.env {
			os.Setenv(key, value)
		}

		result := envservice.Validate()
		if result != test.expected {
			t.Errorf("%v: expected %v, got %v", test.name, test.expected, result)
		}

		for key := range test.env {
			os.Unsetenv(key)
		}
	}
}
