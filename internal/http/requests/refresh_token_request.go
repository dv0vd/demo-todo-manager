package requests

import (
	"demo-todo-manager/internal/enums"
)

func RefreshTokenValidateMethod(method string) bool {
	return method == enums.HttpMethod.Get
}
