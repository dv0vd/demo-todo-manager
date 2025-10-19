package responses

import "demo-todo-manager/internal/dto"

type NoteData struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type NoteResponseStruct struct {
	Success bool     `json:"success"`
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
		},
	}
}
