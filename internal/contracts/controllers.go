package contracts

import (
	"net/http"
)

type AuthController interface {
	GetAuthService() AuthService
	RefreshToken(http.ResponseWriter, *http.Request)
}

type UserController interface {
	Login(http.ResponseWriter, *http.Request)
	Signup(http.ResponseWriter, *http.Request)
}
