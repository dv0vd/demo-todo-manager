package responses

type tokenRefreshData struct {
	Token string `json:"token"`
}

type tokenRefreshResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Data    tokenRefreshData `json:"data"`
}

func NewTokenRefreshResponse(token string) tokenRefreshResponse {
	return tokenRefreshResponse{
		Success: true,
		Message: "Token refreshed successfully",
		Data: tokenRefreshData{
			Token: token,
		},
	}
}
