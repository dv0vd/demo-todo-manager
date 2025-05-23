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

func NewUserSignupResponse(userDTO dto.UserDTO) userSignupResponse {
	return userSignupResponse{
		Success: true,
		Message: "User created successfully",
		Data: userSignupData{
			ID:    userDTO.ID,
			Email: userDTO.Email,
		},
	}
}
