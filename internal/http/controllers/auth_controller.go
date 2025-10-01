package controllers

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/http/responses"
	"net/http"
	"strconv"
)

type authController struct {
	authService contracts.AuthService
}

func NewAuthController(authService contracts.AuthService) contracts.AuthController {
	return &authController{
		authService: authService,
	}
}

func (c *authController) GetAuthService() contracts.AuthService {
	return c.authService
}

func (c *authController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	token, err := c.authService.GetToken(r.Header.Get("Authorization"))
	if err != nil {
		controllerGenerateJsonResponse(
			w,
			r,
			responses.NewErrorResponse("Invalid token"),
			http.StatusBadRequest,
		)

		return
	}

	userId, err := token.Claims.GetSubject()
	if err != nil {
		controllerGenerateJsonResponse(
			w,
			r,
			responses.NewErrorResponse("Invalid token"),
			http.StatusBadRequest,
		)

		return
	}

	userIdConverted, err := strconv.ParseUint(userId, 10, 0)
	if err != nil {
		controllerGenerateJsonResponse(
			w,
			r,
			responses.NewErrorResponse("Invalid token"),
			http.StatusBadRequest,
		)

		return
	}

	newToken, err := c.authService.IssueToken(userIdConverted)
	if err != nil {
		controllerGenerateJsonResponse(
			w,
			r,
			responses.NewErrorResponse("Invalid token"),
			http.StatusBadRequest,
		)

		return
	}

	controllerGenerateJsonResponse(
		w,
		r,
		responses.NewTokenRefreshResponse(newToken),
		http.StatusOK,
	)
}
