package requests

import "demo-todo-manager/internal/enums"

type UndoneNoteRequest struct{}

func UndoneNoteValidateMethod(method string) bool {
	return method == enums.HttpMethod.Get
}
