package middleware

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/http/responses"
	"demo-todo-manager/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func AuthMiddleware(next http.Handler, authService contracts.AuthService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !utils.MiddlewareAuthCheck(authHeader, authService) {
			errorResponse(w, "Invalid token")

			return
		}

		next.ServeHTTP(w, r)
	})
}

func errorResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	err := json.NewEncoder(w).Encode(responses.NewErrorResponse(message))
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to form JSON response. Error: %v", err.Error()), http.StatusUnauthorized)
	}
}
