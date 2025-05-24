package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func MiddlewareAuthCheck(header, secret string, refreshTtl uint64) bool {
	if header == "" {
		return false
	}

	prefix := "Bearer "
	if !strings.HasPrefix(header, prefix) {
		return false
	}

	header = strings.TrimPrefix(header, prefix)
	token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			exp, err := token.Claims.GetExpirationTime()
			if err != nil || time.Since(time.Unix(exp.Unix(), 0)) > time.Duration(refreshTtl)*time.Second {
				return false
			}
		} else {
			return false
		}
	}

	return token.Valid
}

func MiddlewareContentTypeCheck(header string) bool {
	return header == "application/json"
}
