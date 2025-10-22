package controllers

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/enums"
	"demo-todo-manager/internal/http/responses"
	"demo-todo-manager/pkg/localizer"
	"demo-todo-manager/pkg/logger"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/language"
)

type methodValidationFn func(string) bool

func BadRequestResponse(w http.ResponseWriter, r *http.Request, message string) {
	baseErrorResponse(w, r, message, http.StatusBadRequest)
}

func ConflictResponse(w http.ResponseWriter, r *http.Request, message string) {
	JsonResponse(
		w,
		r,
		responses.ErrorResponse(message),
		http.StatusConflict,
	)
}

func GetLocalizer(r *http.Request) *localizer.Localizer {
	loc, ok := r.Context().Value(localizer.GetContextKey()).(*localizer.Localizer)

	if ok && loc != nil {
		return loc
	}

	return localizer.New(language.English)
}

func JsonResponse(w http.ResponseWriter, r *http.Request, res interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to form JSON response. Error: %v", err.Error())
		logger.Log.WithFields(
			logrus.Fields{"method": r.Method, "headers": r.Header, "url": r.URL}).Error(errMsg)

		http.Error(w, errMsg, http.StatusInternalServerError)
	}
}

func InitControllers(
	userService contracts.UserService,
	authService contracts.AuthService,
	noteService contracts.NoteService,
) (contracts.UserController, contracts.AuthController, contracts.NoteController) {
	return newUserController(userService, authService),
		newAuthController(authService),
		newNoteController(authService, userService, noteService)
}

func MethodValidation(w http.ResponseWriter, r *http.Request, vaidationFn methodValidationFn) bool {
	if !vaidationFn(r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return false
	}

	return true
}

func MethodsWithoutBodyCheck(method string) bool {
	if method == enums.HttpMethod.Get || method == enums.HttpMethod.Delete {
		return true
	}

	return false
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request, message string) {
	baseErrorResponse(w, r, message, http.StatusNotFound)
}

func Preparation(w http.ResponseWriter, r *http.Request, req interface{}, validateFn methodValidationFn) bool {
	body := bodyParser(w, r)
	if body == nil {
		return false
	}

	if !MethodValidation(w, r, validateFn) {
		return false
	}

	if !parseJsonRequest(w, r, body, req) {
		return false
	}

	if !validateJsonRequest(w, r, req) {
		return false
	}

	return true
}

func UnauthorizedResponse(w http.ResponseWriter, r *http.Request, message string) {
	JsonResponse(
		w,
		r,
		responses.ErrorResponse(message),
		http.StatusUnauthorized,
	)
}

func UnknownErrorResponse(w http.ResponseWriter, r *http.Request) {
	JsonResponse(
		w,
		r,
		responses.ErrorResponse("Unknown error"),
		http.StatusInternalServerError,
	)
}

func baseErrorResponse(w http.ResponseWriter, r *http.Request, message string, statusCode int) {
	JsonResponse(
		w,
		r,
		responses.ErrorResponse(message),
		statusCode,
	)
}

func bodyParser(w http.ResponseWriter, r *http.Request) []byte {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Log.WithFields(
			logrus.Fields{"method": r.Method, "headers": r.Header, "url": r.URL}).Warning(fmt.Sprintf("Failed to read request body. Error: %v", err.Error()))

		w.WriteHeader(http.StatusBadRequest)

		return nil
	}

	return body
}

func parseJsonRequest(w http.ResponseWriter, r *http.Request, body []byte, req interface{}) bool {
	if MethodsWithoutBodyCheck(r.Method) {
		return true
	}

	err := json.Unmarshal(body, &req)

	if err != nil {
		errMsg := fmt.Sprintf("Unable to parse JSON. Error: %v", err.Error())
		logger.Log.WithFields(
			logrus.Fields{"method": r.Method, "headers": r.Header, "body": string(body), "url": r.URL}).Error(errMsg)

		JsonResponse(
			w,
			r,
			responses.ErrorResponse(errMsg),
			http.StatusBadRequest,
		)

		return false
	}

	return true
}

func validateJsonRequest(w http.ResponseWriter, r *http.Request, req interface{}) bool {
	if MethodsWithoutBodyCheck(r.Method) {
		return true
	}

	validate := validator.New()
	validationErrors := validate.Struct(req)
	if validationErrors != nil {
		var dataErrors []string
		validationErrors, _ = validationErrors.(validator.ValidationErrors)
		for _, err := range validationErrors.(validator.ValidationErrors) {
			dataErrors = append(dataErrors, fmt.Sprintf("Field validation for '%s' failed on the '%s' tag",
				err.Field(), err.Tag()))
		}

		localizer := GetLocalizer(r)

		JsonResponse(
			w,
			r,
			responses.ValidationErrorResponse(dataErrors, localizer.T("common.validation_errors", nil)),
			http.StatusUnprocessableEntity,
		)

		return false
	}

	return true
}
