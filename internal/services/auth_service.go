package services

import (
	"context"
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/pkg/logger"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type authService struct {
	secret     string
	ttl        uint64
	ttlRefresh uint64
}

func NewAuthService(
	secret string,
	ttl,
	ttlRefresh uint64,
) contracts.AuthService {
	return &authService{
		secret:     secret,
		ttl:        ttl,
		ttlRefresh: ttlRefresh,
	}
}

func (s *authService) ExtractEncodedTokenFromHeader(header string) string {
	if header == "" {
		return ""
	}

	prefix := "Bearer "
	if !strings.HasPrefix(header, prefix) {
		return ""
	}

	return strings.TrimPrefix(header, prefix)
}

func (s *authService) GetToken(extractedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(extractedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			exp, err := token.Claims.GetExpirationTime()

			if err != nil {
				return token, err
			}

			if time.Since(time.Unix(exp.Unix(), 0)) < time.Duration(s.ttlRefresh)*time.Second {
				return token, nil
			}
		} else {
			return token, err
		}
	}

	return token, nil
}

func (s *authService) GetUserIdFromContext(ctx context.Context) uint64 {
	userId := ctx.Value(s.GetUserIdContextKey())
	return userId.(uint64)
}

func (s *authService) GetUserIdContextKey() contracts.UserIdContextKey {
	const userIDKey contracts.UserIdContextKey = "userId"

	return userIDKey
}

func (s *authService) IssueToken(userId uint64) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   fmt.Sprint(userId),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s.ttl) * time.Second)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.secret))
	if err != nil {
		logger.Log.WithField("userId", userId).Warningf("Error issuing token for user '%v'. Error: %v", userId, err.Error())

		return "", err
	}

	return signedToken, nil
}
