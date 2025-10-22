package responses

type UserLoginData struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsIn..."`
}

type UserLoginResponseStruct struct {
	Success bool          `json:"success" example:"true"`
	Message string        `json:"message" example:"User logged in successfully"`
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
