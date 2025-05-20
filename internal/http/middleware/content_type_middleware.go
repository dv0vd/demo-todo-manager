package middleware

import (
	"demo-todo-manager/pkg/logger"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			body, _ := io.ReadAll(r.Body)
			errorMsg := "Content-Type must be application/json!"
			logger.Log.WithFields(
				logrus.Fields{"method": r.Method, "headers": r.Header, "body": string(body), "url": "/api/signup"}).Warning(errorMsg)
			http.Error(w, errorMsg, http.StatusUnsupportedMediaType)

			return
		}

		next.ServeHTTP(w, r)
	})
}
