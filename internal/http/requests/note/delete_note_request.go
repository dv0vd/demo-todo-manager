package requests

import "demo-todo-manager/internal/enums"

type DeleteNoteRequest struct{}

func DeleteNoteValidateMethod(method string) bool {
	return method == enums.HttpMethod.Delete
}
