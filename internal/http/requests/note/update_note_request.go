package requests

import "demo-todo-manager/internal/enums"

type UpdateNoteRequest struct {
	Title       string `json:"title" validate:"requried,string"`
	Description string `json:"description" validate:"omitempty,min=1"`
}

func UpdateNoteValidateMethod(method string) bool {
	return method == enums.HttpMethod.Put
}
