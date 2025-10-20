package tests

import (
	"demo-todo-manager/internal/http/middleware"
	"demo-todo-manager/internal/services"
	testutils "demo-todo-manager/internal/tests"
	"fmt"
	"strconv"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestAuthCheckMiddleware(t *testing.T) {
	tests := []struct {
		name                 string
		header               string
		generateHeader       bool
		headerWithSubject    bool
		headerWithExpiration bool
		regenerateEnv        bool
		expected             bool
		env                  map[string]string
	}{
		{
			name:                 "Authorization header is correct",
			header:               "",
			generateHeader:       true,
			headerWithSubject:    true,
			headerWithExpiration: true,
			env: map[string]string{
				"JWT_SECRET":      faker.Word(),
				"JWT_TTL":         testutils.GetRandomInt(100, 1000),
				"JWT_REFRESH_TTL": testutils.GetRandomInt(100, 1000),
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
			name:                 "JWT token is valid, but with incorrect signature",
			header:               "",
			generateHeader:       true,
			headerWithSubject:    true,
			headerWithExpiration: true,
			regenerateEnv:        true,
			env: map[string]string{
				"JWT_SECRET":      faker.Word(),
				"JWT_TTL":         testutils.GetRandomInt(100, 1000),
				"JWT_REFRESH_TTL": testutils.GetRandomInt(100, 1000),
			},
			expected: false,
		},
		{
			name:                 "JWT token is valid, but without expiration",
			header:               "",
			generateHeader:       true,
			headerWithSubject:    true,
			headerWithExpiration: false,
			env: map[string]string{
				"JWT_SECRET":      faker.Word(),
				"JWT_TTL":         testutils.GetRandomInt(100, 1000),
				"JWT_REFRESH_TTL": testutils.GetRandomInt(100, 1000),
			},
			expected: false,
		},
		{
			name:                 "JWT token is expired for access",
			header:               "",
			generateHeader:       true,
			headerWithSubject:    true,
			headerWithExpiration: true,
			env: map[string]string{
				"JWT_SECRET":      faker.Word(),
				"JWT_TTL":         "0",
				"JWT_REFRESH_TTL": testutils.GetRandomInt(100, 1000),
			},
			expected: false,
		},
		{
			name:                 "JWT token is expired for refresh",
			header:               "",
			generateHeader:       true,
			headerWithSubject:    true,
			headerWithExpiration: true,
			env: map[string]string{
				"JWT_SECRET":      faker.Word(),
				"JWT_TTL":         "0",
				"JWT_REFRESH_TTL": "0",
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testutils.SetEnv(test.env)

			header := test.header
			authService := services.InitAuthService()

			if test.generateHeader {
				userId, _ := strconv.ParseUint(testutils.GetRandomInt(1, 1000), 10, 0)
				token, _ := authService.IssueToken(userId, test.headerWithSubject, test.headerWithExpiration)
				header = fmt.Sprintf("Bearer %v", token)
			}

			if test.regenerateEnv {
				testutils.UnsetEnv(test.env)
				testutils.SetEnv(map[string]string{
					"JWT_SECRET":      faker.Word(),
					"JWT_TTL":         "0",
					"JWT_REFRESH_TTL": testutils.GetRandomInt(100, 1000),
				})
				authService = services.InitAuthService()
			}

			result := middleware.AuthCheck(header, authService)
			if result != test.expected {
				t.Errorf("%v: expected %v, got %v", test.name, test.expected, result)
			}

			testutils.UnsetEnv(test.env)
		})
	}
}

func TestTokenClaimsMiddleware(t *testing.T) {
	tests := []struct {
		name                 string
		header               string
		generateHeader       bool
		headerWithSubject    bool
		headerWithExpiration bool
		regenerateEnv        bool
		expected             bool
		env                  map[string]string
	}{
		{
			name:                 "JWT token is valid, but without subject",
			header:               "",
			generateHeader:       true,
			headerWithSubject:    false,
			headerWithExpiration: true,
			env: map[string]string{
				"JWT_SECRET":      faker.Word(),
				"JWT_TTL":         testutils.GetRandomInt(100, 1000),
				"JWT_REFRESH_TTL": testutils.GetRandomInt(100, 1000),
			},
			expected: false,
		},
	}

	for _, test := range tests {
		testutils.SetEnv(test.env)

		header := test.header
		authService := services.InitAuthService()

		if test.generateHeader {
			userId, _ := strconv.ParseUint(testutils.GetRandomInt(1, 1000), 10, 0)
			token, _ := authService.IssueToken(userId, test.headerWithSubject, test.headerWithExpiration)
			header = fmt.Sprintf("Bearer %v", token)
		}

		if test.regenerateEnv {
			testutils.UnsetEnv(test.env)
			testutils.SetEnv(map[string]string{
				"JWT_SECRET":      fmt.Sprintf("%v%v%v", faker.Word(), testutils.GetRandomInt(100, 1000), faker.Word()),
				"JWT_TTL":         "0",
				"JWT_REFRESH_TTL": testutils.GetRandomInt(100, 1000),
			})
			authService = services.InitAuthService()
		}

		_, result := middleware.TokenClaimsCheck(header, authService)
		if result != test.expected {
			t.Errorf("%v: expected %v, got %v", test.name, test.expected, result)
		}

		testutils.UnsetEnv(test.env)
	}
}
