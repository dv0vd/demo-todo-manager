package responses

import "demo-todo-manager/internal/dto"

type userSignupData struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
}

type userSignupResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    userSignupData `json:"data"`
}

func UserSignupResponse(userDTO dto.UserDTO, message string) userSignupResponse {
	return userSignupResponse{
		Success: true,
		Message: message,
		Data: userSignupData{
			ID:    userDTO.ID,
			Email: userDTO.Email,
		},
	}
}
