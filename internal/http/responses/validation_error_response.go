package responses

type ValidationErrorResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}
