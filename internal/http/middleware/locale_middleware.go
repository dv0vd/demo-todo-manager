package middleware

import (
	"context"
	"demo-todo-manager/pkg/localizer"
	"net/http"

	"golang.org/x/text/language"
)

func LocaleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		locale, err := language.Parse(r.Header.Get("Accept-Language"))
		if err != nil {
			locale = language.English
		}

		ctx := context.WithValue(r.Context(), localizer.GetContextKey(), localizer.New(locale))
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
