package controllers

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/dto"
	requests "demo-todo-manager/internal/http/requests/user"
	responses "demo-todo-manager/internal/http/responses/user"
	"demo-todo-manager/internal/utils"
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

	if !utils.ControllerPreparation(w, r, &req, requests.UserLoginValidateMethod) {
		return
	}

	userDTO := dto.UserDTO{
		Email:    req.Email,
		Password: req.Password,
	}

	existedUserDTO, ok := c.userService.GetByEmail(userDTO.Email)
	if !ok {
		utils.ControllerUnknownErrorResponse(w, r)
		return
	}

	if (existedUserDTO == dto.UserDTO{}) {
		utils.ControllerUnauthorizedResponse(w, r, "Incorrect login or password")
		return
	}

	if !c.userService.ValidatePassword(userDTO.Password, existedUserDTO.Password) {
		utils.ControllerUnauthorizedResponse(w, r, "Incorrect login or password")
		return
	}

	token, err := c.authService.IssueToken(existedUserDTO.ID)
	if err != nil {
		utils.ControllerUnknownErrorResponse(w, r)
		return
	}

	utils.ControllerJsonResponse(
		w,
		r,
		responses.UserLoginResponse(token),
		http.StatusOK,
	)
}

func (c *userController) Signup(w http.ResponseWriter, r *http.Request) {
	var req requests.UserSignupRequest

	if !utils.ControllerPreparation(w, r, &req, requests.UserSignupValidateMethod) {
		return
	}

	userDTO := dto.UserDTO{
		Email:    req.Email,
		Password: req.Password,
	}

	existedUserDTO, ok := c.userService.GetByEmail(userDTO.Email)
	if !ok {
		utils.ControllerUnknownErrorResponse(w, r)
		return
	}

	if (existedUserDTO != dto.UserDTO{}) {
		utils.ControllerConflictResponse(w, r, "User already exists")
		return
	}

	insertedUserDTO, err := c.userService.Store(userDTO)
	if err != nil {
		utils.ControllerUnknownErrorResponse(w, r)
		return
	}

	utils.ControllerJsonResponse(
		w,
		r,
		responses.UserSignupResponse(insertedUserDTO),
		http.StatusCreated,
	)
}
