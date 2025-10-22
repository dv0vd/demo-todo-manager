package responses

type ErrorResponseStruct struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Unknown error"`
}

func ErrorResponse(message string) ErrorResponseStruct {
	return ErrorResponseStruct{
		Success: false,
		Message: message,
	}
}
