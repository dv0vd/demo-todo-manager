package middleware

import (
	"demo-todo-manager/internal/utils"
	"net/http"
)

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !utils.MiddlewareContentTypeCheck(r.Header.Get("Content-Type")) {
			http.Error(w, "Content-Type must be application/json!", http.StatusUnsupportedMediaType)

			return
		}

		next.ServeHTTP(w, r)
	})
}
