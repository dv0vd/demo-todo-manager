package responses

import "demo-todo-manager/internal/dto"

type data struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
}

type userSignupResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    data   `json:"data"`
}

func NewUserSignupResponse(userDTO dto.UserDTO) userSignupResponse {
	return userSignupResponse{
		Success: true,
		Message: "User created successfully",
		Data: data{
			ID:    userDTO.ID,
			Email: userDTO.Email,
		},
	}
}
