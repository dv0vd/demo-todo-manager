package responses

import "demo-todo-manager/internal/dto"

type note struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type notesData struct {
	Notes []note `json:"notes"`
}

type notesResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    notesData `json:"data"`
}

func NotesResponse(notes []dto.NoteDTO) notesResponse {
	result := []note{}

	for _, noteDTO := range notes {
		result = append(result, note{
			ID:          noteDTO.ID,
			Title:       noteDTO.Title,
			Description: noteDTO.Description,
		})
	}

	return notesResponse{
		Success: true,
		Data: notesData{
			Notes: result,
		},
	}
}
