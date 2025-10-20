package tests

import (
	"demo-todo-manager/internal/services"
	testutils "demo-todo-manager/internal/tests"
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
				"JWT_TTL":         testutils.GetRandomInt(1, 1000),
				"JWT_REFRESH_TTL": testutils.GetRandomInt(1, 1000),
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
				"JWT_TTL":         testutils.GetRandomInt(1, 1000),
				"JWT_REFRESH_TTL": testutils.GetRandomInt(1, 1000),
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
				"JWT_TTL":         testutils.GetRandomInt(1, 1000),
				"JWT_REFRESH_TTL": testutils.GetRandomInt(1, 1000),
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
				"JWT_TTL":         testutils.GetRandomInt(1, 1000),
				"JWT_REFRESH_TTL": testutils.GetRandomInt(1, 1000),
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
				"JWT_TTL":         testutils.GetRandomInt(1, 1000),
				"JWT_REFRESH_TTL": testutils.GetRandomInt(1, 1000),
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
				"JWT_TTL":         testutils.GetRandomInt(1, 1000),
				"JWT_REFRESH_TTL": testutils.GetRandomInt(1, 1000),
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
				"JWT_REFRESH_TTL": testutils.GetRandomInt(1, 1000),
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
				"JWT_TTL":     testutils.GetRandomInt(1, 1000),
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
				"JWT_TTL":         testutils.GetRandomInt(1, 1000),
				"JWT_REFRESH_TTL": testutils.GetRandomInt(1, 1000),
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
				"JWT_TTL":         testutils.GetRandomInt(-1000, -1),
				"JWT_REFRESH_TTL": testutils.GetRandomInt(1, 1000),
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
				"JWT_TTL":         testutils.GetRandomInt(1, 1000),
				"JWT_REFRESH_TTL": testutils.GetRandomInt(-1000, -1),
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
				"JWT_REFRESH_TTL": testutils.GetRandomInt(1, 1000),
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
				"JWT_TTL":         testutils.GetRandomInt(1, 1000),
				"JWT_REFRESH_TTL": faker.Word(),
				"JWT_SECRET":      faker.Word(),
			},
			expected: false,
		},
	}

	envservice := services.InitEnvService()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testutils.SetEnv(test.env)

			result := envservice.Validate()
			testutils.CheckResult(t, test.name, test.expected, result)

			testutils.UnsetEnv(test.env)
		})
	}
}
