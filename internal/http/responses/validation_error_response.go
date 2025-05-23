package responses

type validationErrorResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

func NewValidationErrorResponse(data []string) validationErrorResponse {
	return validationErrorResponse{
		Success: false,
		Message: "Validation errors",
		Data:    data,
	}
}
