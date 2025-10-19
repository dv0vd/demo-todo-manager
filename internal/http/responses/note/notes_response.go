package responses

import "demo-todo-manager/internal/dto"

type Note struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type NotesData struct {
	Notes []Note `json:"notes"`
}

type NotesResponseStruct struct {
	Success bool      `json:"success"`
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
		})
	}

	return NotesResponseStruct{
		Success: true,
		Data: NotesData{
			Notes: result,
		},
	}
}
