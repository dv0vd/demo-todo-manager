package responses

import "demo-todo-manager/internal/dto"

type UserSignupData struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
}

type UserSignupResponseStruct struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    UserSignupData `json:"data"`
}

func UserSignupResponse(userDTO dto.UserDTO, message string) UserSignupResponseStruct {
	return UserSignupResponseStruct{
		Success: true,
		Message: message,
		Data: UserSignupData{
			ID:    userDTO.ID,
			Email: userDTO.Email,
		},
	}
}
