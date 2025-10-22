package requests

import (
	"demo-todo-manager/internal/enums"
)

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"example@email.com"`
	Password string `json:"password" validate:"required,min=8,max=36" example:"secret-password"`
}

func UserLoginValidateMethod(method string) bool {
	return method == enums.HttpMethod.Post
}
