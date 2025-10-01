package tests

import (
	"demo-todo-manager/internal/utils"
	"fmt"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestAuthServiceExtractEncodedTokenFromHeader(t *testing.T) {
	prefix := "Bearer "
	authService := utils.ServicesInitAuthService()
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
		result := authService.ExtractEncodedTokenFromHeader(test.header)
		utils.TestCheckResult(t, test.name, test.expected, result)
	}
}
