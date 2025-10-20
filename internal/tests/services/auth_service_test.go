package tests

import (
	"context"
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/enums"
	"demo-todo-manager/internal/services"
	testutils "demo-todo-manager/internal/tests"
	"errors"
	"fmt"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/golang-jwt/jwt/v5"
)

func TestAuthServiceExtractEncodedTokenFromHeader(t *testing.T) {
	prefix := "Bearer "
	authService := services.InitAuthService()
	encodedToken := faker.Word()
	fullHeader := fmt.Sprintf("%v%v", prefix, encodedToken)

	tests := []struct {
		name     string
		header   string
		expected string
	}{
		{
			name:     "Header is empty",
			header:   "",
			expected: "",
		},
		{
			name:     fmt.Sprintf("Header doesn't contain '%v' prefix", prefix),
			header:   faker.Word(),
			expected: "",
		},
		{
			name:     "Correct header",
			header:   fullHeader,
			expected: encodedToken,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := authService.ExtractEncodedTokenFromHeader(test.header)
			testutils.CheckResult(t, test.name, test.expected, result)
		})
	}
}

func TestGetToken(t *testing.T) {
	env := map[string]string{
		"JWT_SECRET":      faker.Word(),
		"JWT_TTL":         testutils.GetRandomInt(100, 1000),
		"JWT_REFRESH_TTL": testutils.GetRandomInt(100, 1000),
	}
	testutils.SetEnv(env)
	authService := services.InitAuthService()

	id := testutils.GetRandomInt(100, 1000)
	userId, _ := strconv.ParseUint(id, 10, 64)

	tokenString, _ := authService.IssueToken(userId, true, true)
	token, err := authService.GetToken(tokenString)
	userIdStringFromContext, err := token.Claims.GetSubject()

	testName := "No error during parsing token with all fields"
	t.Run(testName, func(t *testing.T) {
		testutils.CheckResult(t, testName, err, nil)
	})
	testName = "No error during getting subject from token"
	t.Run(testName, func(t *testing.T) {
		testutils.CheckResult(t, testName, err, nil)
	})
	testName = "Token subject is correct"
	t.Run(testName, func(t *testing.T) {
		userIdFromContext, _ := strconv.ParseUint(userIdStringFromContext, 10, 64)
		testutils.CheckResult(t, testName, userIdFromContext, userId)
	})
	testName = "No error during getting expiration from token"
	t.Run(testName, func(t *testing.T) {
		_, err = token.Claims.GetExpirationTime()
		testutils.CheckResult(t, testName, err, nil)
	})

	tokenString, _ = authService.IssueToken(userId, false, true)
	token, err = authService.GetToken(tokenString)

	testName = "No error during parsing token without subject"
	t.Run(testName, func(t *testing.T) {
		testutils.CheckResult(t, testName, err, nil)
	})
	testName = "No subject in token without subject"
	t.Run(testName, func(t *testing.T) {
		subject, err := token.Claims.GetSubject()
		testutils.CheckResult(t, testName, err, nil)
		testutils.CheckResult(t, testName, subject, "")
	})
	testName = "Error during parsing token without expiration"
	t.Run(testName, func(t *testing.T) {
		tokenString, _ = authService.IssueToken(userId, true, false)
		token, err = authService.GetToken(tokenString)
		testutils.CheckResult(t, testName, token, (*jwt.Token)(nil))
		testutils.CheckResult(t, testName, errors.Is(err, jwt.ErrTokenMalformed), true)
	})
	testName = "Error parsing token with incorrect signature"
	t.Run(testName, func(t *testing.T) {
		tokenString, _ = authService.IssueToken(userId, true, false)

		testutils.UnsetEnv(env)

		env = map[string]string{
			"JWT_SECRET":      fmt.Sprintf("%v%v%v", faker.Word(), testutils.GetRandomInt(100, 1000), faker.Word()),
			"JWT_TTL":         testutils.GetRandomInt(100, 1000),
			"JWT_REFRESH_TTL": testutils.GetRandomInt(100, 1000),
		}
		testutils.SetEnv(env)
		authService = services.InitAuthService()
		token, err = authService.GetToken(tokenString)
		testutils.CheckResult(t, testName, token, (*jwt.Token)(nil))
		testutils.CheckResult(t, testName, errors.Is(err, jwt.ErrSignatureInvalid), true)
	})

	testutils.UnsetEnv(env)
}

func TestGetUserIdContextKey(t *testing.T) {
	authService := services.InitAuthService()
	const userIDKey contracts.UserIdContextKey = "userId"

	testutils.CheckResult(t, "User ID context key", authService.GetUserIdContextKey(), userIDKey)
}

func TestGetUserIdFromContext(t *testing.T) {
	authService := services.InitAuthService()
	userId, _ := strconv.ParseUint(testutils.GetRandomInt(100, 1000), 10, 64)
	r := httptest.NewRequest(enums.HttpMethod.Get, "/", nil)
	ctx := context.WithValue(r.Context(), authService.GetUserIdContextKey(), userId)

	testutils.CheckResult(t, "User ID from context", authService.GetUserIdFromContext(ctx), userId)
}

func TestIssueToken(t *testing.T) {
	env := map[string]string{
		"JWT_SECRET":      faker.Word(),
		"JWT_TTL":         testutils.GetRandomInt(100, 1000),
		"JWT_REFRESH_TTL": testutils.GetRandomInt(100, 1000),
	}
	testutils.SetEnv(env)
	authService := services.InitAuthService()

	id := testutils.GetRandomInt(100, 1000)
	userId, _ := strconv.ParseUint(id, 10, 64)

	_, err := authService.IssueToken(userId, true, true)
	testutils.CheckResult(t, "No error during token issue with all fields", err, nil)

	_, err = authService.IssueToken(userId, false, true)
	testutils.CheckResult(t, "No error during token issue without subject", err, nil)

	_, err = authService.IssueToken(userId, true, false)
	testutils.CheckResult(t, "No error during token issue without expiration", err, nil)

	testutils.UnsetEnv(env)
}
