package requests

import "demo-todo-manager/internal/enums"

type NotesRequest struct{}

func GetNotesValidateMethod(method string) bool {
	return method == enums.HttpMethod.Get
}
