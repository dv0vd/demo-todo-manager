package controllers

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/dto"
	"demo-todo-manager/internal/http/requests"
	"demo-todo-manager/internal/http/responses"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func (c *noteController) Delete(w http.ResponseWriter, r *http.Request) {
	var req requests.DeleteNoteRequest

	if !ControllerPreparation(w, r, &req, requests.DeleteNoteValidateMethod) {
		return
	}

	noteId, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 64)
	if err != nil || noteId <= 0 {
		controllerGenerateJsonResponse(
			w,
			r,
			responses.NewErrorResponse("Incorrect note id"),
			http.StatusBadRequest,
		)

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

	note, ok := c.noteService.Get(noteId, userId)
	if !ok {
		controllerGenerateJsonResponse(
			w,
			r,
			responses.NewErrorResponse("Unknown error"),
			http.StatusInternalServerError,
		)

		return
	}

	if (note == dto.NoteDTO{}) {
		controllerGenerateJsonResponse(
			w,
			r,
			responses.NewErrorResponse("Note not found"),
			http.StatusNotFound,
		)

		return
	}

	ok = c.noteService.Delete(noteId, userId)
	if !ok {
		controllerGenerateJsonResponse(
			w,
			r,
			responses.NewErrorResponse("Unknown error"),
			http.StatusInternalServerError,
		)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *noteController) GetAll(w http.ResponseWriter, r *http.Request) {
	var req requests.GetNotesRequest

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

		return
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
