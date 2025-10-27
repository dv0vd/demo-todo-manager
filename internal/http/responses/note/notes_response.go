package responses

import "demo-todo-manager/internal/dto"

type Note struct {
	ID          uint64 `json:"id" example:"161"`
	Title       string `json:"title" example:"For Peace"`
	Description string `json:"description" example:"May all people live safely and without suffering."`
	Done        bool   `json:"done" example:"true"`
}

type NotesData struct {
	Notes []Note `json:"notes"`
}

type NotesResponseStruct struct {
	Success bool      `json:"success" example:"true"`
	Message string    `json:"message"`
	Data    NotesData `json:"data"`
}

func NotesResponse(notes []dto.NoteDTO) NotesResponseStruct {
	result := []Note{}

	for _, noteDTO := range notes {
		result = append(result, Note{
			ID:          noteDTO.ID,
			Title:       noteDTO.Title,
			Description: noteDTO.Description,
			Done:        noteDTO.Done,
		})
	}

	return NotesResponseStruct{
		Success: true,
		Data: NotesData{
			Notes: result,
		},
	}
}
