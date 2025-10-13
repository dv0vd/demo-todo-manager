package middleware

import (
	"context"
	"demo-todo-manager/internal/http/responses"
	"demo-todo-manager/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func AuthMiddleware(next http.Handler) http.Handler {
	authService := utils.ServicesInitAuthService()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !utils.MiddlewareAuthCheck(authHeader, authService) {
			errorResponse(w, "Invalid token")
			return
		}

		token, _ := authService.GetToken(authService.ExtractEncodedTokenFromHeader(authHeader))

		subject, err := token.Claims.GetSubject()
		if err != nil {
			errorResponse(w, "Invalid token")
			return
		}
		userId, err := strconv.ParseUint(subject, 10, 64)
		if err != nil {
			errorResponse(w, "Invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), authService.GetUserIdKey(), userId)
		r = r.WithContext(ctx)
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
