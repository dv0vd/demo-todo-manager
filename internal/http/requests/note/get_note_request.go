package requests

import "demo-todo-manager/internal/enums"

type GetNoteRequest struct{}

func GetNoteValidateMethod(method string) bool {
	return method == enums.HttpMethod.Get
}
