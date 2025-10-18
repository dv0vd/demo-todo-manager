package middleware

import (
	"context"
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/http/controllers"
	"demo-todo-manager/internal/utils"
	"net/http"
	"strconv"
)

func AuthCheck(header string, authService contracts.AuthService) bool {
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

func AuthMiddleware(next http.Handler) http.Handler {
	authService := utils.ServicesInitAuthService()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !AuthCheck(authHeader, authService) {
			controllers.UnauthorizedResponse(w, r, "Invalid token")
			return
		}

		token, _ := authService.GetToken(authService.ExtractEncodedTokenFromHeader(authHeader))

		subject, err := token.Claims.GetSubject()
		if err != nil {
			controllers.UnauthorizedResponse(w, r, "Invalid token")
			return
		}
		userId, err := strconv.ParseUint(subject, 10, 64)
		if err != nil {
			controllers.UnauthorizedResponse(w, r, "Invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), authService.GetUserIdContextKey(), userId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
