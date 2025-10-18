package middleware

import (
	"net/http"
)

func ContentTypeCheck(header string) bool {
	return header == "application/json"
}

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !ContentTypeCheck(r.Header.Get("Content-Type")) {
			http.Error(w, "Content-Type must be application/json!", http.StatusUnsupportedMediaType)

			return
		}

		next.ServeHTTP(w, r)
	})
}
