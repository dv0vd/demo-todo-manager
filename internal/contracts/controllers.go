package contracts

import "net/http"

type UserController interface {
	Login(http.ResponseWriter, *http.Request)
	Signup(http.ResponseWriter, *http.Request)
}
