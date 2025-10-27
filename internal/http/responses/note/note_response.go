package responses

import "demo-todo-manager/internal/dto"

type NoteData struct {
	ID          uint64 `json:"id" example:"161"`
	Title       string `json:"title" example:"For Peace"`
	Description string `json:"description" example:"May all people live safely and without suffering."`
	Done        bool   `json:"done" example:"true"`
}

type NoteResponseStruct struct {
	Success bool     `json:"success" example:"true"`
	Message string   `json:"message"`
	Data    NoteData `json:"data"`
}

func NoteResponse(note dto.NoteDTO) NoteResponseStruct {
	return NoteResponseStruct{
		Success: true,
		Data: NoteData{
			ID:          note.ID,
			Title:       note.Title,
			Description: note.Description,
			Done:        note.Done,
		},
	}
}
