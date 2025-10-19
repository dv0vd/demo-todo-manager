package tests

import (
	responses "demo-todo-manager/internal/http/responses/auth"
	testutils "demo-todo-manager/internal/tests"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestTokenRefreshResponse(t *testing.T) {
	token := faker.Word()
	message := faker.Word()

	testutils.CheckResult(
		t,
		"Token response",
		responses.TokenRefreshResponse(token, message),
		responses.TokenRefreshResponseStruct{
			Success: true,
			Message: message,
			Data: responses.TokenRefreshData{
				Token: token,
			},
		},
	)
}
