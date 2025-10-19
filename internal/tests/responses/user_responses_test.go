package tests

import (
	"demo-todo-manager/internal/dto"
	responses "demo-todo-manager/internal/http/responses/user"
	testutils "demo-todo-manager/internal/tests"
	"strconv"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestUserLoginResponse(t *testing.T) {
	token := faker.Word()
	message := faker.Word()

	testutils.CheckResult(
		t,
		"User login response",
		responses.UserLoginResponse(token, message),
		responses.UserLoginResponseStruct{
			Success: true,
			Message: message,
			Data: responses.UserLoginData{
				Token: token,
			},
		},
	)
}

func TestUserSignupResponse(t *testing.T) {
	message := faker.Word()
	id, _ := strconv.ParseUint(testutils.GetRandomInt(100, 1000), 10, 64)
	email := faker.Email()
	userDTO := dto.UserDTO{
		ID:    id,
		Email: email,
	}

	testutils.CheckResult(
		t,
		"User signup response",
		responses.UserSignupResponse(userDTO, message),
		responses.UserSignupResponseStruct{
			Success: true,
			Message: message,
			Data: responses.UserSignupData{
				ID:    userDTO.ID,
				Email: userDTO.Email,
			},
		},
	)
}
