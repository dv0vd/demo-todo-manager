package responses

type TokenRefreshData struct {
	Token string `json:"token"`
}

type TokenRefreshResponseStruct struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Data    TokenRefreshData `json:"data"`
}

func TokenRefreshResponse(token, message string) TokenRefreshResponseStruct {
	return TokenRefreshResponseStruct{
		Success: true,
		Message: message,
		Data: TokenRefreshData{
			Token: token,
		},
	}
}
