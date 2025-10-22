package responses

type TokenRefreshData struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsIn..."`
}

type TokenRefreshResponseStruct struct {
	Success bool             `json:"success" example:"true"`
	Message string           `json:"message" example:"Token refreshed successfully"`
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
