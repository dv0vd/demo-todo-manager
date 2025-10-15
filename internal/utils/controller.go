package utils

import (
	"demo-todo-manager/internal/enums"
	"demo-todo-manager/internal/http/responses"
	"demo-todo-manager/pkg/logger"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type methodValidationFn func(string) bool

func ControllerJsonResponse(w http.ResponseWriter, r *http.Request, res interface{}, statusCode int) {
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

func ControllerBadRequestResponse(w http.ResponseWriter, r *http.Request, message string) {
	controllerBaseResponse(w, r, message, http.StatusBadRequest)
}

func ControllerConflictResponse(w http.ResponseWriter, r *http.Request, message string) {
	ControllerJsonResponse(
		w,
		r,
		responses.ErrorResponse(message),
		http.StatusConflict,
	)
}

func ControllerMethodValidation(w http.ResponseWriter, r *http.Request, vaidationFn methodValidationFn) bool {
	if !vaidationFn(r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return false
	}

	return true
}

func ControllerNotFoundResponse(w http.ResponseWriter, r *http.Request, message string) {
	controllerBaseResponse(w, r, message, http.StatusNotFound)
}

func ControllerPreparation(w http.ResponseWriter, r *http.Request, req interface{}, validateFn methodValidationFn) bool {
	body := contollerBodyParser(w, r)
	if body == nil {
		return false
	}

	if !ControllerMethodValidation(w, r, validateFn) {
		return false
	}

	if !controllerParseJsonRequest(w, r, body, req) {
		return false
	}

	if !controllerValidateJsonRequest(w, r, req) {
		return false
	}

	return true
}

func ControllerUnauthorizedResponse(w http.ResponseWriter, r *http.Request, message string) {
	ControllerJsonResponse(
		w,
		r,
		responses.ErrorResponse(message),
		http.StatusUnauthorized,
	)
}

func ControllerUnknownErrorResponse(w http.ResponseWriter, r *http.Request) {
	ControllerJsonResponse(
		w,
		r,
		responses.ErrorResponse("Unknown error"),
		http.StatusInternalServerError,
	)
}

func controllerBaseResponse(w http.ResponseWriter, r *http.Request, message string, statusCode int) {
	ControllerJsonResponse(
		w,
		r,
		responses.ErrorResponse(message),
		statusCode,
	)
}

func contollerBodyParser(w http.ResponseWriter, r *http.Request) []byte {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Log.WithFields(
			logrus.Fields{"method": r.Method, "headers": r.Header, "url": r.URL}).Warning(fmt.Sprintf("Failed to read request body. Error: %v", err.Error()))

		w.WriteHeader(http.StatusBadRequest)

		return nil
	}

	return body
}

func controllerParseJsonRequest(w http.ResponseWriter, r *http.Request, body []byte, req interface{}) bool {
	if r.Method == enums.HttpMethod.Get || r.Method == enums.HttpMethod.Delete {
		return true
	}

	err := json.Unmarshal(body, &req)

	if err != nil {
		errMsg := fmt.Sprintf("Unable to parse JSON. Error: %v", err.Error())
		logger.Log.WithFields(
			logrus.Fields{"method": r.Method, "headers": r.Header, "body": string(body), "url": r.URL}).Error(errMsg)

		ControllerJsonResponse(
			w,
			r,
			responses.ErrorResponse(errMsg),
			http.StatusBadRequest,
		)

		return false
	}

	return true
}

func controllerValidateJsonRequest(w http.ResponseWriter, r *http.Request, req interface{}) bool {
	validate := validator.New()
	validationErrors := validate.Struct(req)
	if validationErrors != nil {
		var dataErrors []string
		validationErrors, _ = validationErrors.(validator.ValidationErrors)
		for _, err := range validationErrors.(validator.ValidationErrors) {
			dataErrors = append(dataErrors, fmt.Sprintf("Field validation for '%s' failed on the '%s' tag",
				err.Field(), err.Tag()))
		}

		ControllerJsonResponse(
			w,
			r,
			responses.ValidationErrorResponse(dataErrors),
			http.StatusUnprocessableEntity,
		)

		return false
	}

	return true
}
