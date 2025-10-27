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

func newNoteController(authService contracts.AuthService, userService contracts.UserService, noteService contracts.NoteService) contracts.NoteController {
	return &noteController{
		userService: userService,
		authService: authService,
		noteService: noteService,
	}
}

// @Summary Delete note
// @Description Deletes the note identified by its ID. This action cannot be undone
// @Tags Note
// @Param Accept-Language header string false "User locale" Enums(ru, en) Example(en) Default(en)
// @Param id path int true "Note ID"
// @Success 204
// @Failure 400 {object} responses.ErrorResponseStruct
// @Failure 401 {object} responses.ErrorResponseStruct
// @Failure 404 {object} responses.ErrorResponseStruct
// @Failure 500 {object} responses.ErrorResponseStruct
// @Router /notes/{id} [delete]
// @Security ApiKeyAuth
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

// @Summary Update note
// @Description Fully updates an existing note by its ID. All fields will be overwritten with the provided values.
// @Tags Note
// @Accept json
// @Produce json
// @Param Accept-Language header string false "User locale" Enums(ru, en) Example(en) Default(en)
// @Param id path int true "Note ID"
// @Param request body requests.UpdateNoteRequest true "Note data"
// @Success 200 {object} responses.NoteResponseStruct
// @Failure 400 {object} responses.ErrorResponseStruct
// @Failure 401 {object} responses.ErrorResponseStruct
// @Failure 404 {object} responses.ErrorResponseStruct
// @Failure 415 {string} string "Content-Type must be application/json!"
// @Failure 500 {object} responses.ErrorResponseStruct
// @Router /notes/{id} [put]
// @Security ApiKeyAuth
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

// @Summary Get all notes
// @Description Retrieves a list of all notes belonging to the authenticated user.
// @Tags Note
// @Accept json
// @Produce json
// @Param Accept-Language header string false "User locale" Enums(ru, en) Example(en) Default(en)
// @Param done query boolean false "Filter notes by status: true — completed, false — not completed" Example(true) Default(en)
// @Success 200 {object} responses.NotesResponseStruct
// @Failure 400 {object} responses.ErrorResponseStruct
// @Failure 401 {object} responses.ErrorResponseStruct
// @Failure 415 {string} string "Content-Type must be application/json!"
// @Failure 500 {object} responses.ErrorResponseStruct
// @Router /notes [get]
// @Security ApiKeyAuth
func (c *noteController) Index(w http.ResponseWriter, r *http.Request) {
	var req requests.GetNotesRequest

	if !Preparation(w, r, &req, requests.GetNotesValidateMethod) {
		return
	}

	userDTO, ok := c.getUser(w, r)
	if !ok {
		return
	}

	whereClauses := map[string]interface{}{}
	done := r.URL.Query().Get("done")
	if done != "" {
		where, err := strconv.ParseBool(done)
		if err == nil {
			whereClauses["done"] = where
		}
	}

	notes, ok := c.noteService.GetByUserId(userDTO.ID, whereClauses)
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

// @Summary Get note by ID
// @Description Returns the note details by its unique ID. Requires authentication.
// @Tags Note
// @Accept json
// @Produce json
// @Param Accept-Language header string false "User locale" Enums(ru, en) Example(en) Default(en)
// @Param id path int true "Note ID"
// @Success 200 {object} responses.NoteResponseStruct
// @Failure 400 {object} responses.ErrorResponseStruct
// @Failure 401 {object} responses.ErrorResponseStruct
// @Failure 404 {object} responses.ErrorResponseStruct
// @Failure 500 {object} responses.ErrorResponseStruct
// @Router /notes/{id} [get]
// @Security ApiKeyAuth
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

// @Summary Create a new note
// @Description Creates a new note with the provided title and content. Returns the created note.
// @Tags Note
// @Accept json
// @Produce json
// @Param Accept-Language header string false "User locale" Enums(ru, en) Example(en) Default(en)
// @Param request body requests.StoreNoteRequest true "Note data"
// @Success 200 {object} responses.NoteResponseStruct
// @Failure 400 {object} responses.ErrorResponseStruct
// @Failure 401 {object} responses.ErrorResponseStruct
// @Failure 415 {string} string "Content-Type must be application/json!"
// @Failure 500 {object} responses.ErrorResponseStruct
// @Router /notes [post]
// @Security ApiKeyAuth
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
