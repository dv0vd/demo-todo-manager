package tests

import (
	"demo-todo-manager/internal/services"
	"demo-todo-manager/internal/utils"
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
				"DB_HOST":         "localhost",
				"DB_PORT":         "5432",
				"DB_NAME":         "name",
				"DB_USER":         "user",
				"DB_PASSWORD":     "password",
				"JWT_TTL":         utils.TestGetRandomInt(1, 1000),
				"JWT_REFRESH_TTL": utils.TestGetRandomInt(1, 1000),
				"JWT_SECRET":      faker.Word(),
			},
			expected: true,
		},
		{
			name: "DB_HOST env variable is not set",
			env: map[string]string{
				"DB_PORT":         "5432",
				"DB_NAME":         "name",
				"DB_USER":         "user",
				"DB_PASSWORD":     "password",
				"JWT_TTL":         utils.TestGetRandomInt(1, 1000),
				"JWT_REFRESH_TTL": utils.TestGetRandomInt(1, 1000),
				"JWT_SECRET":      faker.Word(),
			},
			expected: false,
		},
		{
			name: "DB_PORT env variable is not set",
			env: map[string]string{
				"DB_HOST":         "localhost",
				"DB_NAME":         "name",
				"DB_USER":         "user",
				"DB_PASSWORD":     "password",
				"JWT_TTL":         utils.TestGetRandomInt(1, 1000),
				"JWT_REFRESH_TTL": utils.TestGetRandomInt(1, 1000),
				"JWT_SECRET":      faker.Word(),
			},
			expected: false,
		},
		{
			name: "DB_NAME env variable is not set",
			env: map[string]string{
				"DB_HOST":         "localhost",
				"DB_PORT":         "5432",
				"DB_USER":         "user",
				"DB_PASSWORD":     "password",
				"JWT_TTL":         utils.TestGetRandomInt(1, 1000),
				"JWT_REFRESH_TTL": utils.TestGetRandomInt(1, 1000),
				"JWT_SECRET":      faker.Word(),
			},
			expected: false,
		},
		{
			name: "DB_USER env variable is not set",
			env: map[string]string{
				"DB_HOST":         "localhost",
				"DB_PORT":         "5432",
				"DB_NAME":         "name",
				"DB_PASSWORD":     "password",
				"JWT_TTL":         utils.TestGetRandomInt(1, 1000),
				"JWT_REFRESH_TTL": utils.TestGetRandomInt(1, 1000),
				"JWT_SECRET":      faker.Word(),
			},
			expected: false,
		},
		{
			name: "DB_PASSWORD env variable is not set",
			env: map[string]string{
				"DB_HOST":         "localhost",
				"DB_PORT":         "5432",
				"DB_NAME":         "name",
				"JWT_TTL":         utils.TestGetRandomInt(1, 1000),
				"JWT_REFRESH_TTL": utils.TestGetRandomInt(1, 1000),
				"JWT_SECRET":      faker.Word(),
			},
			expected: false,
		},
		{
			name: "JWT_TTL env variable is not set",
			env: map[string]string{
				"DB_HOST":         "localhost",
				"DB_PORT":         "5432",
				"DB_NAME":         "name",
				"DB_USER":         "user",
				"DB_PASSWORD":     "password",
				"JWT_REFRESH_TTL": utils.TestGetRandomInt(1, 1000),
				"JWT_SECRET":      faker.Word(),
			},
			expected: false,
		},
		{
			name: "JWT_REFRESH_TTL env variable is not set",
			env: map[string]string{
				"DB_HOST":     "localhost",
				"DB_PORT":     "5432",
				"DB_NAME":     "name",
				"DB_USER":     "user",
				"DB_PASSWORD": "password",
				"JWT_TTL":     utils.TestGetRandomInt(1, 1000),
				"JWT_SECRET":  faker.Word(),
			},
			expected: false,
		},
		{
			name: "JWT_SECRET env variable is not set",
			env: map[string]string{
				"DB_HOST":         "localhost",
				"DB_PORT":         "5432",
				"DB_NAME":         "name",
				"DB_USER":         "user",
				"DB_PASSWORD":     "password",
				"JWT_TTL":         utils.TestGetRandomInt(1, 1000),
				"JWT_REFRESH_TTL": utils.TestGetRandomInt(1, 1000),
			},
			expected: false,
		},
		{
			name: "JWT_TTL env variable is negative",
			env: map[string]string{
				"DB_HOST":         "localhost",
				"DB_PORT":         "5432",
				"DB_NAME":         "name",
				"DB_USER":         "user",
				"DB_PASSWORD":     "password",
				"JWT_TTL":         utils.TestGetRandomInt(-1000, -1),
				"JWT_REFRESH_TTL": utils.TestGetRandomInt(1, 1000),
				"JWT_SECRET":      faker.Word(),
			},
			expected: false,
		},
		{
			name: "JWT_REFRESH_TTL env variable is negative",
			env: map[string]string{
				"DB_HOST":         "localhost",
				"DB_PORT":         "5432",
				"DB_NAME":         "name",
				"DB_USER":         "user",
				"DB_PASSWORD":     "password",
				"JWT_TTL":         utils.TestGetRandomInt(1, 1000),
				"JWT_REFRESH_TTL": utils.TestGetRandomInt(-1000, -1),
				"JWT_SECRET":      faker.Word(),
			},
			expected: false,
		},
		{
			name: "JWT_TTL env variable is string",
			env: map[string]string{
				"DB_HOST":         "localhost",
				"DB_PORT":         "5432",
				"DB_NAME":         "name",
				"DB_USER":         "user",
				"DB_PASSWORD":     "password",
				"JWT_TTL":         faker.Word(),
				"JWT_REFRESH_TTL": utils.TestGetRandomInt(1, 1000),
				"JWT_SECRET":      faker.Word(),
			},
			expected: false,
		},
		{
			name: "JWT_REFRESH_TTL env variable is string",
			env: map[string]string{
				"DB_HOST":         "localhost",
				"DB_PORT":         "5432",
				"DB_NAME":         "name",
				"DB_USER":         "user",
				"DB_PASSWORD":     "password",
				"JWT_TTL":         utils.TestGetRandomInt(1, 1000),
				"JWT_REFRESH_TTL": faker.Word(),
				"JWT_SECRET":      faker.Word(),
			},
			expected: false,
		},
	}

	envservice := services.NewEnvService()

	for _, test := range tests {
		utils.TestSetEnv(test.env)

		result := envservice.Validate()
		utils.TestCheckResult(t, test.name, test.expected, result)

		utils.TestUnsetEnv(test.env)
	}
}
