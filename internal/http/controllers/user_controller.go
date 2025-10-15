package controllers

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/dto"
	requests "demo-todo-manager/internal/http/requests/user"
	baseResponses "demo-todo-manager/internal/http/responses"
	responses "demo-todo-manager/internal/http/responses/user"
	"net/http"
)

type userController struct {
	userService contracts.UserService
	authService contracts.AuthService
}

func NewUserController(userService contracts.UserService, authService contracts.AuthService) contracts.UserController {
	return &userController{
		userService: userService,
		authService: authService,
	}
}

func (c *userController) Login(w http.ResponseWriter, r *http.Request) {
	var req requests.UserLoginRequest

	if !ControllerPreparation(w, r, &req, requests.UserLoginValidateMethod) {
		return
	}

	userDTO := dto.UserDTO{
		Email:    req.Email,
		Password: req.Password,
	}

	existedUserDTO, ok := c.userService.GetByEmail(userDTO.Email)
	if !ok {
		controllerGenerateUnknownErrorResponse(w, r)
		return
	}

	if (existedUserDTO == dto.UserDTO{}) {
		controllerGenerateJsonResponse(
			w,
			r,
			baseResponses.ErrorResponse("Incorrect login or password"),
			http.StatusOK,
		)

		return
	}

	if !c.userService.ValidatePassword(userDTO.Password, existedUserDTO.Password) {
		controllerGenerateJsonResponse(
			w,
			r,
			baseResponses.ErrorResponse("Incorrect login or password"),
			http.StatusOK,
		)

		return
	}

	token, err := c.authService.IssueToken(existedUserDTO.ID)
	if err != nil {
		controllerGenerateUnknownErrorResponse(w, r)
		return
	}

	controllerGenerateJsonResponse(
		w,
		r,
		responses.UserLoginResponse(token),
		http.StatusOK,
	)
}

func (c *userController) Signup(w http.ResponseWriter, r *http.Request) {
	var req requests.UserSignupRequest

	if !ControllerPreparation(w, r, &req, requests.UserSignupValidateMethod) {
		return
	}

	userDTO := dto.UserDTO{
		Email:    req.Email,
		Password: req.Password,
	}

	existedUserDTO, ok := c.userService.GetByEmail(userDTO.Email)
	if !ok {
		controllerGenerateUnknownErrorResponse(w, r)
		return
	}

	if (existedUserDTO != dto.UserDTO{}) {
		controllerGenerateJsonResponse(
			w,
			r,
			baseResponses.ErrorResponse("User already exists"),
			http.StatusConflict,
		)

		return
	}

	insertedUserDTO, err := c.userService.Store(userDTO)
	if err != nil {
		controllerGenerateUnknownErrorResponse(w, r)
		return
	}

	controllerGenerateJsonResponse(
		w,
		r,
		responses.UserSignupResponse(insertedUserDTO),
		http.StatusCreated,
	)
}
