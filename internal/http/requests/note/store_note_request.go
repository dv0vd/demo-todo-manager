package requests

import "demo-todo-manager/internal/enums"

type StoreNoteRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Description string `json:"description" validate:"omitempty,min=1"`
}

func StoreNoteValidateMethod(method string) bool {
	return method == enums.HttpMethod.Post
}
