package responses

type UserSignupResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"data"`
}
