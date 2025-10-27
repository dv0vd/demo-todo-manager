package requests

import "demo-todo-manager/internal/enums"

type DoneNoteRequest struct{}

func DoneNoteValidateMethod(method string) bool {
	return method == enums.HttpMethod.Get
}
