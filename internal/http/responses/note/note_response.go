package responses

import "demo-todo-manager/internal/dto"

type noteData struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type noteResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    noteData `json:"data"`
}

func NoteResponse(note dto.NoteDTO) noteResponse {
	return noteResponse{
		Success: true,
		Data: noteData{
			ID:          note.ID,
			Title:       note.Title,
			Description: note.Description,
		},
	}
}
