package responses

import "demo-todo-manager/internal/dto"

type UserSignupData struct {
	ID    uint64 `json:"id" example:"1618"`
	Email string `json:"email" example:"example@email.com"`
}

type UserSignupResponseStruct struct {
	Success bool           `json:"success" example:"true"`
	Message string         `json:"message" example:"User 'example@email.com' created successfully"`
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
