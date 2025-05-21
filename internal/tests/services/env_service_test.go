package tests

import (
	"demo-todo-manager/internal/services"
	"os"
	"testing"
)

func TestEnvServiceValidate(t *testing.T) {
	tests := []struct {
		name     string
		envVars  map[string]string
		expected bool
	}{
		{
			name: "All env variables are set",
			envVars: map[string]string{
				"TODO_MANAGER_DB_HOST": "localhost",
				"TODO_MANAGER_DB_PORT": "5432",
				"TODO_MANAGER_DB_NAME": "name",
				"TODO_MANAGER_DB_USER": "user",
			},
			expected: true,
		},
		{
			name: "TODO_MANAGER_DB_HOST env variable is not set",
			envVars: map[string]string{
				"TODO_MANAGER_DB_PORT": "5432",
				"TODO_MANAGER_DB_NAME": "name",
				"TODO_MANAGER_DB_USER": "user",
			},
			expected: false,
		},
		{
			name: "TODO_MANAGER_DB_PORT env variable is not set",
			envVars: map[string]string{
				"TODO_MANAGER_DB_HOST": "localhost",
				"TODO_MANAGER_DB_NAME": "name",
				"TODO_MANAGER_DB_USER": "user",
			},
			expected: false,
		},
		{
			name: "TODO_MANAGER_DB_NAME env variable is not set",
			envVars: map[string]string{
				"TODO_MANAGER_DB_HOST": "localhost",
				"TODO_MANAGER_DB_PORT": "5432",
				"TODO_MANAGER_DB_USER": "user",
			},
			expected: false,
		},
		{
			name: "TODO_MANAGER_DB_USER env variable is not set",
			envVars: map[string]string{
				"TODO_MANAGER_DB_HOST": "localhost",
				"TODO_MANAGER_DB_PORT": "5432",
				"TODO_MANAGER_DB_NAME": "name",
			},
			expected: false,
		},
	}

	envservice := services.NewEnvService()

	for _, test := range tests {
		for key, value := range test.envVars {
			os.Setenv(key, value)
		}

		result := envservice.Validate()
		if result != test.expected {
			t.Errorf("%v: expected %v, got %v", test.name, test.expected, result)
		}

		for key := range test.envVars {
			os.Unsetenv(key)
		}
	}
}
