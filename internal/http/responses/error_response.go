package responses

type ErrorResponseStruct struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ErrorResponse(message string) ErrorResponseStruct {
	return ErrorResponseStruct{
		Success: false,
		Message: message,
	}
}
