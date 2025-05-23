package services

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/pkg/logger"
	"fmt"
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

func (s *authService) IssueToken(userId uint64) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   fmt.Sprint(userId),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s.ttl))),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.secret))
	if err != nil {
		logger.Log.Warningf("Error issuing token for user '%v'. Error: %v", userId, err.Error())

		return "", err
	}

	return signedToken, nil
}
