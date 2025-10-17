package controllers

import (
	"demo-todo-manager/internal/contracts"
	requests "demo-todo-manager/internal/http/requests/auth"
	responses "demo-todo-manager/internal/http/responses/auth"
	"demo-todo-manager/internal/utils"
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
	if !utils.ControllerMethodValidation(w, r, requests.RefreshTokenValidateMethod) {
		return
	}

	localizer := utils.ControllerGetLocalizer(r)

	token, err := c.authService.GetToken(c.authService.ExtractEncodedTokenFromHeader(r.Header.Get("Authorization")))
	if err != nil {
		utils.ControllerBadRequestResponse(w, r, localizer.T("auth.invalid_token", nil))
		return
	}

	userId, err := token.Claims.GetSubject()
	if err != nil {
		utils.ControllerBadRequestResponse(w, r, localizer.T("auth.invalid_token", nil))
		return
	}

	userIdConverted, err := strconv.ParseUint(userId, 10, 0)
	if err != nil {
		utils.ControllerBadRequestResponse(w, r, localizer.T("auth.invalid_token", nil))
		return
	}

	newToken, err := c.authService.IssueToken(userIdConverted)
	if err != nil {
		utils.ControllerBadRequestResponse(w, r, localizer.T("auth.invalid_token", nil))
		return
	}

	utils.ControllerJsonResponse(
		w,
		r,
		responses.TokenRefreshResponse(newToken),
		http.StatusOK,
	)
}
