package responses

type ValidationErrorResponseStruct struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

func ValidationErrorResponse(data []string, message string) ValidationErrorResponseStruct {
	return ValidationErrorResponseStruct{
		Success: false,
		Message: message,
		Data:    data,
	}
}
