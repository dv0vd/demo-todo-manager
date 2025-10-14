package requests

import "demo-todo-manager/internal/enums"

type GetNotesRequest struct{}

func GetNotesValidateMethod(method string) bool {
	return method == enums.HttpMethod.Get
}
