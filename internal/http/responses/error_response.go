package responses

type errorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ErrorResponse(message string) errorResponse {
	return errorResponse{
		Success: false,
		Message: message,
	}
}
