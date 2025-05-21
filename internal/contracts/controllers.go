package contracts

import "net/http"

type UserController interface {
	Signup(http.ResponseWriter, *http.Request)
}
