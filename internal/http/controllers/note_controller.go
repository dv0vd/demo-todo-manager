package controllers

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/http/requests"
	"demo-todo-manager/internal/http/responses"
	"net/http"
)

type noteController struct {
	userService contracts.UserService
	authService contracts.AuthService
	noteService contracts.NoteService
}

func NewNoteController(authService contracts.AuthService, userService contracts.UserService, noteService contracts.NoteService) contracts.NoteController {
	return &noteController{
		userService: userService,
		authService: authService,
		noteService: noteService,
	}
}

func (c *noteController) GetAll(w http.ResponseWriter, r *http.Request) {
	var req requests.NotesRequest

	if !ControllerPreparation(w, r, &req, requests.GetNotesValidateMethod) {
		return
	}

	userId := c.authService.GetUserIdFromContext(r.Context())
	_, ok := c.userService.GetById(userId)
	if !ok {
		controllerGenerateJsonResponse(
			w,
			r,
			responses.NewErrorResponse("User not found"),
			http.StatusBadRequest,
		)

		return
	}

	notes, ok := c.noteService.GetByUserId(userId)
	if !ok {
		controllerGenerateJsonResponse(
			w,
			r,
			responses.NewErrorResponse("Unknown error"),
			http.StatusInternalServerError,
		)
	}

	controllerGenerateJsonResponse(
		w,
		r,
		responses.NewNotesResponse(notes),
		http.StatusOK,
	)
}

func (c *noteController) GetNoteService() contracts.NoteService {
	return c.noteService
}
