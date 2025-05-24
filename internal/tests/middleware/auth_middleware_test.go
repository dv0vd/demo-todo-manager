package tests

import (
	"demo-todo-manager/internal/utils"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestAuthMiddleware(t *testing.T) {
	tests := []struct {
		name           string
		header         string
		generateHeader bool
		expected       bool
		env            map[string]string
	}{
		{
			name:           "Authorization header is correct",
			header:         "",
			generateHeader: true,
			env: map[string]string{
				"JWT_SECRET":      faker.Word(),
				"JWT_TTL":         utils.TestGetRandomInt(100, 1000),
				"JWT_REFRESH_TTL": utils.TestGetRandomInt(100, 1000),
			},
			expected: true,
		},
		{
			name:           "Authorization header is empty",
			header:         "",
			generateHeader: false,
			env:            map[string]string{},
			expected:       false,
		},
		{
			name:           "Authorization header has no 'Bearer ' prefix",
			header:         faker.Word(),
			generateHeader: false,
			env:            map[string]string{},
			expected:       false,
		},
		{
			name:           "JWT token is invalid",
			header:         fmt.Sprintf("Bearer %v", faker.Word()),
			generateHeader: false,
			env:            map[string]string{},
			expected:       false,
		},
		{
			name:           "JWT token is expired",
			header:         "",
			generateHeader: true,
			env: map[string]string{
				"JWT_SECRET":      faker.Word(),
				"JWT_TTL":         "0",
				"JWT_REFRESH_TTL": utils.TestGetRandomInt(100, 1000),
			},
			expected: false,
		},
	}

	for _, test := range tests {
		for key, value := range test.env {
			os.Setenv(key, value)
		}

		header := test.header
		authService := utils.ServicesInitAuthService()

		if test.generateHeader {
			userId, _ := strconv.ParseUint(utils.TestGetRandomInt(1, 1000), 10, 0)
			correctToken, _ := authService.IssueToken(userId)
			header = fmt.Sprintf("Bearer %v", correctToken)
		}

		result := utils.MiddlewareAuthCheck(header, authService.GetSecret(), authService.GetRefreshTTL())
		if result != test.expected {
			t.Errorf("%v: expected %v, got %v", test.name, test.expected, result)
		}

		for key, _ := range test.env {
			os.Unsetenv(key)
		}
	}
}
