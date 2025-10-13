package contracts

import (
	"net/http"
)

type AuthController interface {
	GetAuthService() AuthService
	RefreshToken(http.ResponseWriter, *http.Request)
}

type NoteController interface {
	Delete(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	GetNoteService() NoteService
}

type UserController interface {
	Login(http.ResponseWriter, *http.Request)
	Signup(http.ResponseWriter, *http.Request)
}
