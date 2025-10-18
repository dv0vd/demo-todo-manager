package controllers

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/dto"
	requests "demo-todo-manager/internal/http/requests/note"
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

	if !Preparation(w, r, &req, requests.DeleteNoteValidateMethod) {
		return
	}

	noteId, ok := c.getNoteIdFromURL(w, r)
	if !ok {
		return
	}

	userDTO, ok := c.getUser(w, r)
	if !ok {
		return
	}

	note, ok := c.noteService.Get(noteId, userDTO.ID)
	if !ok {
		UnknownErrorResponse(w, r)
		return
	}

	localizer := GetLocalizer(r)

	if (note == dto.NoteDTO{}) {
		NotFoundResponse(w, r, localizer.T("note.not_found", map[string]interface{}{
			"id": noteId,
		}))
		return
	}

	ok = c.noteService.Delete(noteId, userDTO.ID)
	if !ok {
		UnknownErrorResponse(w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *noteController) Edit(w http.ResponseWriter, r *http.Request) {
	var req requests.UpdateNoteRequest

	if !Preparation(w, r, &req, requests.UpdateNoteValidateMethod) {
		return
	}

	noteId, ok := c.getNoteIdFromURL(w, r)
	if !ok {
		return
	}

	userDTO, ok := c.getUser(w, r)
	if !ok {
		return
	}

	noteDTO, ok := c.noteService.Get(noteId, userDTO.ID)
	if !ok {
		UnknownErrorResponse(w, r)
		return
	}

	localizer := GetLocalizer(r)

	if (noteDTO == dto.NoteDTO{}) {
		NotFoundResponse(w, r, localizer.T("note.not_found", map[string]interface{}{
			"id": noteId,
		}))
		return
	}

	noteDTO.Title = req.Title
	noteDTO.Description = req.Description
	if !c.noteService.Update(noteDTO, userDTO.ID) {
		UnknownErrorResponse(w, r)
		return
	}

	JsonResponse(
		w,
		r,
		responses.NoteResponse(noteDTO),
		http.StatusOK,
	)
}

func (c *noteController) GetNoteService() contracts.NoteService {
	return c.noteService
}

func (c *noteController) Index(w http.ResponseWriter, r *http.Request) {
	var req requests.GetNotesRequest

	if !Preparation(w, r, &req, requests.GetNotesValidateMethod) {
		return
	}

	userDTO, ok := c.getUser(w, r)
	if !ok {
		return
	}

	notes, ok := c.noteService.GetByUserId(userDTO.ID)
	if !ok {
		UnknownErrorResponse(w, r)
		return
	}

	JsonResponse(
		w,
		r,
		responses.NotesResponse(notes),
		http.StatusOK,
	)
}

func (c *noteController) Show(w http.ResponseWriter, r *http.Request) {
	var req requests.GetNoteRequest

	if !Preparation(w, r, &req, requests.GetNoteValidateMethod) {
		return
	}

	noteId, ok := c.getNoteIdFromURL(w, r)
	if !ok {
		return
	}

	userDTO, ok := c.getUser(w, r)
	if !ok {
		return
	}

	noteDTO, ok := c.noteService.Get(noteId, userDTO.ID)
	if !ok {
		UnknownErrorResponse(w, r)
		return
	}

	localizer := GetLocalizer(r)

	if (noteDTO == dto.NoteDTO{}) {
		NotFoundResponse(w, r, localizer.T("note.not_found", map[string]interface{}{
			"id": noteId,
		}))
		return
	}

	JsonResponse(
		w,
		r,
		responses.NoteResponse(noteDTO),
		http.StatusOK,
	)
}

func (c *noteController) Store(w http.ResponseWriter, r *http.Request) {
	var req requests.StoreNoteRequest

	if !Preparation(w, r, &req, requests.StoreNoteValidateMethod) {
		return
	}

	userDTO, ok := c.getUser(w, r)
	if !ok {
		return
	}

	noteDTO := dto.NoteDTO{
		Title:       req.Title,
		Description: req.Description,
		UserId:      userDTO.ID,
	}
	noteDTO, err := c.noteService.Create(noteDTO, userDTO.ID)
	if err != nil {
		UnknownErrorResponse(w, r)
		return
	}

	JsonResponse(
		w,
		r,
		responses.NoteResponse(noteDTO),
		http.StatusOK,
	)
}

func (c *noteController) getNoteIdFromURL(w http.ResponseWriter, r *http.Request) (uint64, bool) {
	noteId, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 64)
	localizer := GetLocalizer(r)

	if err != nil || noteId <= 0 {
		BadRequestResponse(w, r, localizer.T("note.incorrect_id", nil))
		return 0, false
	}

	return noteId, true
}

func (c *noteController) getUser(w http.ResponseWriter, r *http.Request) (dto.UserDTO, bool) {
	userId := c.authService.GetUserIdFromContext(r.Context())
	localizer := GetLocalizer(r)

	if userId <= 0 {
		BadRequestResponse(w, r, localizer.T("auth.user_not_found", nil))
		return dto.UserDTO{}, false
	}

	userDTO, ok := c.userService.GetById(userId)
	if (!ok || userDTO == dto.UserDTO{}) {
		BadRequestResponse(w, r, localizer.T("auth.user_not_found", nil))
		return dto.UserDTO{}, false
	}

	return userDTO, true
}
