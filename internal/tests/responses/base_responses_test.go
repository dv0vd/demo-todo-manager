package tests

import (
	"demo-todo-manager/internal/http/responses"
	testutils "demo-todo-manager/internal/tests"
	"strconv"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestErrorResponse(t *testing.T) {
	testName := "Error response"
	t.Run(testName, func(t *testing.T) {
		message := faker.Word()

		testutils.CheckResult(
			t,
			testName,
			responses.ErrorResponse(message),
			responses.ErrorResponseStruct{
				Success: false,
				Message: message,
			},
		)
	})
}

func TestValidationErrorResponse(t *testing.T) {
	testName := "Validation rrror response"
	t.Run(testName, func(t *testing.T) {
		message := faker.Word()
		count, _ := strconv.ParseUint(testutils.GetRandomInt(5, 10), 10, 64)
		messages := []string{}

		for i := uint64(0); i < count; i++ {
			messages = append(messages, faker.Sentence())
		}

		testutils.CheckResult(
			t,
			testName,
			responses.ValidationErrorResponse(messages, message),
			responses.ValidationErrorResponseStruct{
				Success: false,
				Message: message,
				Data:    messages,
			},
		)
	})

}
