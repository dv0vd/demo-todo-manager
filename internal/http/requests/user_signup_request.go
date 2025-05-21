package requests

import (
	"demo-todo-manager/internal/enums"
)

type UserSignupRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func UserSignupValidateMethod(method string) bool {
	return method == enums.HttpMethod.Post
}
