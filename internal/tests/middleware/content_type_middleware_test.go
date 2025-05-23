package tests

import (
	"demo-todo-manager/internal/utils"
	"fmt"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestContentTypeMiddleware(t *testing.T) {
	correctHeader := "application/json"
	tests := []struct {
		name     string
		header   string
		expected bool
	}{
		{
			name:     fmt.Sprintf("Content type is '%v'", correctHeader),
			header:   correctHeader,
			expected: true,
		},
		{
			name:     fmt.Sprintf("Content type is not '%v'", correctHeader),
			header:   faker.Word(),
			expected: false,
		},
	}

	for _, test := range tests {
		result := utils.MiddlewareContentTypeCheck(test.header)
		if result != test.expected {
			t.Errorf("%v: expected %v, got %v", test.name, test.expected, result)
		}
	}
}
