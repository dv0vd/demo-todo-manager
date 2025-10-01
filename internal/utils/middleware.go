package utils

import (
	"demo-todo-manager/internal/contracts"
)

func MiddlewareAuthCheck(header string, authService contracts.AuthService) bool {
	header = authService.ExtractEncodedTokenFromHeader(header)
	if header == "" {
		return false
	}

	token, err := authService.GetToken(header)
	if err != nil {
		return false
	}

	return token.Valid
}

func MiddlewareContentTypeCheck(header string) bool {
	return header == "application/json"
}
