package contracts

import (
	"net/http"
)

type UserController interface {
	GetAuthService() AuthService
	Login(http.ResponseWriter, *http.Request)
	Signup(http.ResponseWriter, *http.Request)
}
