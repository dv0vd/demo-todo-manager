package responses

type userLoginData struct {
	Token string `json:"token"`
}

type userLoginResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    userLoginData `json:"data"`
}

func UserLoginResponse(token, message string) userLoginResponse {
	return userLoginResponse{
		Success: true,
		Message: message,
		Data: userLoginData{
			Token: token,
		},
	}
}
