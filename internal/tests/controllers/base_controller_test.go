package tests

import (
	"context"
	"demo-todo-manager/internal/http/controllers"
	authRequests "demo-todo-manager/internal/http/requests/auth"
	notesRequests "demo-todo-manager/internal/http/requests/note"
	userRequests "demo-todo-manager/internal/http/requests/user"
	"demo-todo-manager/internal/http/responses"
	testutils "demo-todo-manager/internal/tests"
	"demo-todo-manager/pkg/localizer"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-faker/faker/v4"
	"golang.org/x/text/language"
)

func TestBadRequestResponse(t *testing.T) {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	message := faker.Sentence()

	controllers.BadRequestResponse(recorder, req, message)

	testName := "Status code is correct"
	t.Run(testName, func(t *testing.T) {
		testutils.CheckResult(t, testName, recorder.Code, http.StatusBadRequest)
	})

	testName = "Content type is correct"
	t.Run(testName, func(t *testing.T) {
		testutils.CheckResult(t, testName, recorder.Header().Get("Content-Type"), "application/json")
	})

	testName = "Message is correct"
	t.Run(testName, func(t *testing.T) {
		var response responses.ErrorResponseStruct

		err := json.NewDecoder(recorder.Body).Decode(&response)

		testutils.CheckResult(t, testName, err, nil)
		testutils.CheckResult(t, testName, response.Message, message)
	})
}

func TestConflictResponse(t *testing.T) {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	message := faker.Sentence()

	controllers.ConflictResponse(recorder, req, message)

	testName := "Status code is correct"
	t.Run(testName, func(t *testing.T) {
		testutils.CheckResult(t, testName, recorder.Code, http.StatusConflict)
	})

	testName = "Content type is correct"
	t.Run(testName, func(t *testing.T) {
		testutils.CheckResult(t, testName, recorder.Header().Get("Content-Type"), "application/json")
	})

	testName = "Message is correct"
	t.Run(testName, func(t *testing.T) {
		var response responses.ErrorResponseStruct

		err := json.NewDecoder(recorder.Body).Decode(&response)

		testutils.CheckResult(t, testName, err, nil)
		testutils.CheckResult(t, testName, response.Message, message)
	})
}

func TestGetLocalizer(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := context.WithValue(req.Context(), localizer.GetContextKey(), localizer.New(language.Russian))
	req = req.WithContext(ctx)

	localizer := controllers.GetLocalizer(req)
	message := localizer.T("common.validation_errors", nil)
	testName := "Correct locale from context localizer"
	testutils.CheckResult(t, testName, message, "Ошибки валидации")

	req = httptest.NewRequest(http.MethodGet, "/", nil)
	localizer = controllers.GetLocalizer(req)
	message = localizer.T("common.validation_errors", nil)
	testName = "Default locale from context localizer"
	testutils.CheckResult(t, testName, message, "Validation errors")
}

func TestMethodValidation(t *testing.T) {
	recorder := httptest.NewRecorder()

	tests := []struct {
		name       string
		method     string
		validateFn func(string) bool
		expected   bool
	}{
		{
			name:       "Method is correct",
			method:     http.MethodGet,
			validateFn: authRequests.RefreshTokenValidateMethod,
			expected:   true,
		},
		{
			name:       "Method is not correct",
			method:     http.MethodDelete,
			validateFn: authRequests.RefreshTokenValidateMethod,
			expected:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(test.method, "/", nil)
			result := controllers.MethodValidation(recorder, req, test.validateFn)

			testutils.CheckResult(t, test.name, result, test.expected)

			if test.expected == false {
				testutils.CheckResult(t, test.name, recorder.Code, http.StatusMethodNotAllowed)
			}
		})
	}
}

func TestMethodsWithoutBodyCheck(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		expected bool
	}{
		{
			name:     "GET method",
			method:   http.MethodGet,
			expected: true,
		},
		{
			name:     "DELETE method",
			method:   http.MethodDelete,
			expected: true,
		},
		{
			name:     "POST method",
			method:   http.MethodPost,
			expected: false,
		},
		{
			name:     "PATCH method",
			method:   http.MethodPost,
			expected: false,
		},
		{
			name:     "PUT method",
			method:   http.MethodPut,
			expected: false,
		},
		{
			name:     "OPTIONS method",
			method:   http.MethodOptions,
			expected: false,
		},
		{
			name:     "HEAD method",
			method:   http.MethodHead,
			expected: false,
		},
		{
			name:     "CONNECT method",
			method:   http.MethodConnect,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testutils.CheckResult(t, test.name, controllers.MethodsWithoutBodyCheck(test.method), test.expected)
		})
	}
}

func TestNotFoundResponse(t *testing.T) {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	message := faker.Sentence()

	controllers.NotFoundResponse(recorder, req, message)

	testName := "Status code is correct"
	t.Run(testName, func(t *testing.T) {
		testutils.CheckResult(t, testName, recorder.Code, http.StatusNotFound)
	})

	testName = "Content type is correct"
	t.Run(testName, func(t *testing.T) {
		testutils.CheckResult(t, testName, recorder.Header().Get("Content-Type"), "application/json")
	})

	testName = "Message is correct"
	t.Run(testName, func(t *testing.T) {
		var response responses.ErrorResponseStruct

		err := json.NewDecoder(recorder.Body).Decode(&response)

		testutils.CheckResult(t, testName, err, nil)
		testutils.CheckResult(t, testName, response.Message, message)
	})
}

func TestPreparation(t *testing.T) {
	recorder := httptest.NewRecorder()

	tests := []struct {
		name               string
		body               io.Reader
		method             string
		methodValidationFn func(string) bool
		status             uint
		expected           bool
	}{
		{
			name:               "Body is nil",
			body:               nil,
			method:             http.MethodPost,
			methodValidationFn: userRequests.UserLoginValidateMethod,
			status:             http.StatusBadRequest,
			expected:           false,
		},
		{
			name:               "Error if body is empty string",
			body:               strings.NewReader(""),
			method:             http.MethodPost,
			methodValidationFn: userRequests.UserLoginValidateMethod,
			status:             http.StatusBadRequest,
			expected:           false,
		},
		{
			name:               "No JSON parsing and validation if method is GET",
			body:               strings.NewReader(fmt.Sprintf("{%v}", faker.Word())),
			method:             http.MethodGet,
			methodValidationFn: notesRequests.GetNoteValidateMethod,
			status:             http.StatusBadRequest,
			expected:           true,
		},
		{
			name:               "No JSON parsing and validation if method is DELETE",
			body:               strings.NewReader(fmt.Sprintf("{%v}", faker.Word())),
			method:             http.MethodGet,
			methodValidationFn: notesRequests.GetNoteValidateMethod,
			status:             http.StatusBadRequest,
			expected:           true,
		},
		{
			name:               "Error if body is incorrect JSON",
			body:               strings.NewReader(fmt.Sprintf("{%v}", faker.Word())),
			method:             http.MethodPost,
			methodValidationFn: userRequests.UserLoginValidateMethod,
			status:             http.StatusBadRequest,
			expected:           false,
		},
		{
			name:               "JSON validation error",
			body:               strings.NewReader("{}"),
			method:             http.MethodPost,
			methodValidationFn: userRequests.UserLoginValidateMethod,
			status:             http.StatusUnprocessableEntity,
			expected:           false,
		},
		{
			name:               "Correct JSON",
			body:               strings.NewReader(`{"email": "example@email.com", "password": "secret-password"}`),
			method:             http.MethodPost,
			methodValidationFn: userRequests.UserLoginValidateMethod,
			expected:           true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(test.method, "/", test.body)
			var req userRequests.UserLoginRequest
			result := controllers.Preparation(recorder, r, &req, test.methodValidationFn)

			testutils.CheckResult(t, test.name, result, test.expected)

			if test.expected == false {
				testutils.CheckResult(t, test.name, recorder.Code, http.StatusBadRequest)
			}
		})
	}
}

func TestUnauthorizeResponse(t *testing.T) {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	message := faker.Sentence()

	controllers.UnauthorizedResponse(recorder, req, message)

	testName := "Status code is correct"
	t.Run(testName, func(t *testing.T) {
		testutils.CheckResult(t, testName, recorder.Code, http.StatusUnauthorized)
	})

	testName = "Content type is correct"
	t.Run(testName, func(t *testing.T) {
		testutils.CheckResult(t, testName, recorder.Header().Get("Content-Type"), "application/json")
	})

	testName = "Message is correct"
	t.Run(testName, func(t *testing.T) {
		var response responses.ErrorResponseStruct

		err := json.NewDecoder(recorder.Body).Decode(&response)

		testutils.CheckResult(t, testName, err, nil)
		testutils.CheckResult(t, testName, response.Message, message)
	})
}

func TestUnknownErrorResponse(t *testing.T) {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	controllers.UnknownErrorResponse(recorder, req)

	testName := "Status code is correct"
	t.Run(testName, func(t *testing.T) {
		testutils.CheckResult(t, testName, recorder.Code, http.StatusInternalServerError)
	})

	testName = "Content type is correct"
	t.Run(testName, func(t *testing.T) {
		testutils.CheckResult(t, testName, recorder.Header().Get("Content-Type"), "application/json")
	})

	testName = "Message is correct"
	t.Run(testName, func(t *testing.T) {
		var response responses.ErrorResponseStruct

		err := json.NewDecoder(recorder.Body).Decode(&response)

		testutils.CheckResult(t, testName, err, nil)
		testutils.CheckResult(t, testName, response.Message, "Unknown error")
	})
}
