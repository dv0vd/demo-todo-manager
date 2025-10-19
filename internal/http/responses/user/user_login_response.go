package responses

type UserLoginData struct {
	Token string `json:"token"`
}

type UserLoginResponseStruct struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    UserLoginData `json:"data"`
}

func UserLoginResponse(token, message string) UserLoginResponseStruct {
	return UserLoginResponseStruct{
		Success: true,
		Message: message,
		Data: UserLoginData{
			Token: token,
		},
	}
}
