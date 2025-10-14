package controllers

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/dto"
	requests "demo-todo-manager/internal/http/requests/note"
	baseResponses "demo-todo-manager/internal/http/responses"
	responses "demo-todo-manager/internal/http/responses/note"
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
			baseResponses.NewErrorResponse("Incorrect note id"),
			http.StatusBadRequest,
		)

		return
	}

	userId := c.authService.GetUserIdFromContext(r.Context())
	userDTO, ok := c.userService.GetById(userId)
	if (!ok || userId == 0 || userDTO == dto.UserDTO{}) {
		controllerGenerateJsonResponse(
			w,
			r,
			baseResponses.NewErrorResponse("User not found"),
			http.StatusBadRequest,
		)

		return
	}

	note, ok := c.noteService.Get(noteId, userId)
	if !ok {
		controllerGenerateJsonResponse(
			w,
			r,
			baseResponses.NewErrorResponse("Unknown error"),
			http.StatusInternalServerError,
		)

		return
	}

	if (note == dto.NoteDTO{}) {
		controllerGenerateJsonResponse(
			w,
			r,
			baseResponses.NewErrorResponse("Note not found"),
			http.StatusNotFound,
		)

		return
	}

	ok = c.noteService.Delete(noteId, userId)
	if !ok {
		controllerGenerateJsonResponse(
			w,
			r,
			baseResponses.NewErrorResponse("Unknown error"),
			http.StatusInternalServerError,
		)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *noteController) Edit(w http.ResponseWriter, r *http.Request) {
	var req requests.UpdateNoteRequest

	if !ControllerPreparation(w, r, &req, requests.UpdateNoteValidateMethod) {
		return
	}

	noteId, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 64)
	if err != nil || noteId <= 0 {
		controllerGenerateJsonResponse(
			w,
			r,
			baseResponses.NewErrorResponse("Incorrect note id"),
			http.StatusBadRequest,
		)

		return
	}

	userId := c.authService.GetUserIdFromContext(r.Context())
	userDTO, ok := c.userService.GetById(userId)
	if (!ok || userId == 0 || userDTO == dto.UserDTO{}) {
		controllerGenerateJsonResponse(
			w,
			r,
			baseResponses.NewErrorResponse("User not found"),
			http.StatusBadRequest,
		)

		return
	}

	noteDTO, ok := c.noteService.Get(noteId, userId)
	if !ok {
		controllerGenerateJsonResponse(
			w,
			r,
			baseResponses.NewErrorResponse("Unknown error"),
			http.StatusInternalServerError,
		)

		return
	}

	if (noteDTO == dto.NoteDTO{}) {
		controllerGenerateJsonResponse(
			w,
			r,
			baseResponses.NewErrorResponse("Note not found"),
			http.StatusBadRequest,
		)

		return
	}

	noteDTO.Title = req.Title
	noteDTO.Description = req.Description
	if !c.noteService.Update(noteDTO, userId) {
		controllerGenerateJsonResponse(
			w,
			r,
			baseResponses.NewErrorResponse("Unknown error"),
			http.StatusInternalServerError,
		)

		return
	}

	controllerGenerateJsonResponse(
		w,
		r,
		responses.NewNoteResponse(noteDTO),
		http.StatusOK,
	)
}

func (c *noteController) GetNoteService() contracts.NoteService {
	return c.noteService
}

func (c *noteController) Index(w http.ResponseWriter, r *http.Request) {
	var req requests.GetNotesRequest

	if !ControllerPreparation(w, r, &req, requests.GetNotesValidateMethod) {
		return
	}

	userId := c.authService.GetUserIdFromContext(r.Context())
	_, ok := c.userService.GetById(userId)
	if !ok || userId == 0 {
		controllerGenerateJsonResponse(
			w,
			r,
			baseResponses.NewErrorResponse("User not found"),
			http.StatusBadRequest,
		)

		return
	}

	notes, ok := c.noteService.GetByUserId(userId)
	if !ok {
		controllerGenerateJsonResponse(
			w,
			r,
			baseResponses.NewErrorResponse("Unknown error"),
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

func (c *noteController) Show(w http.ResponseWriter, r *http.Request) {
	var req requests.GetNoteRequest

	if !ControllerPreparation(w, r, &req, requests.GetNoteValidateMethod) {
		return
	}

	noteId, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 64)
	if err != nil || noteId <= 0 {
		controllerGenerateJsonResponse(
			w,
			r,
			baseResponses.NewErrorResponse("Incorrect note id"),
			http.StatusBadRequest,
		)

		return
	}

	userId := c.authService.GetUserIdFromContext(r.Context())
	userDTO, ok := c.userService.GetById(userId)
	if (!ok || userId == 0 || userDTO == dto.UserDTO{}) {
		controllerGenerateJsonResponse(
			w,
			r,
			baseResponses.NewErrorResponse("User not found"),
			http.StatusBadRequest,
		)

		return
	}

	noteDTO, ok := c.noteService.Get(noteId, userId)
	if !ok {
		controllerGenerateJsonResponse(
			w,
			r,
			baseResponses.NewErrorResponse("Unknown error"),
			http.StatusInternalServerError,
		)

		return
	}

	if (noteDTO == dto.NoteDTO{}) {
		controllerGenerateJsonResponse(
			w,
			r,
			baseResponses.NewErrorResponse("Note not found"),
			http.StatusNotFound,
		)

		return
	}

	controllerGenerateJsonResponse(
		w,
		r,
		responses.NewNoteResponse(noteDTO),
		http.StatusOK,
	)
}

func (c *noteController) Store(w http.ResponseWriter, r *http.Request) {
	var req requests.StoreNoteRequest

	if !ControllerPreparation(w, r, &req, requests.StoreNoteValidateMethod) {
		return
	}

	userId := c.authService.GetUserIdFromContext(r.Context())
	userDTO, ok := c.userService.GetById(userId)
	if (!ok || userId == 0 || userDTO == dto.UserDTO{}) {
		controllerGenerateJsonResponse(
			w,
			r,
			baseResponses.NewErrorResponse("User not found"),
			http.StatusBadRequest,
		)

		return
	}

	noteDTO := dto.NoteDTO{
		Title:       req.Title,
		Description: req.Description,
		UserId:      userId,
	}
	noteDTO, err := c.noteService.Create(noteDTO, userId)
	if err != nil {
		controllerGenerateJsonResponse(
			w,
			r,
			baseResponses.NewErrorResponse("Unknown error"),
			http.StatusInternalServerError,
		)

		return
	}

	controllerGenerateJsonResponse(
		w,
		r,
		responses.NewNoteResponse(noteDTO),
		http.StatusOK,
	)
}
