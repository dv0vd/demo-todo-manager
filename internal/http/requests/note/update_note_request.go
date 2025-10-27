package requests

import "demo-todo-manager/internal/enums"

type UpdateNoteRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=255" example:"For Peace"`
	Description string `json:"description" validate:"omitempty,min=1" example:"May all people live safely and without suffering."`
	Done        bool   `json:"done" validate:"omitempty" example:"true"`
}

func UpdateNoteValidateMethod(method string) bool {
	return method == enums.HttpMethod.Put
}
