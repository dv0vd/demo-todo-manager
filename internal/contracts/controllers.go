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
	Edit(http.ResponseWriter, *http.Request)
	Index(http.ResponseWriter, *http.Request)
	GetNoteService() NoteService
	Show(http.ResponseWriter, *http.Request)
	Store(http.ResponseWriter, *http.Request)
}

type UserController interface {
	Login(http.ResponseWriter, *http.Request)
	Signup(http.ResponseWriter, *http.Request)
}
