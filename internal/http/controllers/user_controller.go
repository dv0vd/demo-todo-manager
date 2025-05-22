package controllers

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/dto"
	"demo-todo-manager/internal/http/requests"
	"demo-todo-manager/internal/http/responses"
	"net/http"
)

type userController struct {
	userService contracts.UserService
}

func NewUserController(userService contracts.UserService) contracts.UserController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) Signup(w http.ResponseWriter, r *http.Request) {
	var req requests.UserSignupRequest
	url := "/api/signup"

	if !ControllerPreparation(w, r, url, &req, requests.UserSignupValidateMethod) {
		return
	}

	userDTO := dto.UserDTO{
		Email:    req.Email,
		Password: req.Password,
	}

	existedUserDTO, ok := c.userService.GetByEmail(userDTO.Email)
	if !ok {
		controllerGenerateJsonResponse(
			w,
			r,
			url,
			responses.NewErrorResponse("Unknown error"),
			http.StatusInternalServerError,
		)

		return
	}

	if (existedUserDTO != dto.UserDTO{}) {
		controllerGenerateJsonResponse(
			w,
			r,
			url,
			responses.NewErrorResponse("User already exists"),
			http.StatusConflict,
		)

		return
	}

	insertedUserDTO, err := c.userService.Store(userDTO)
	if err != nil {
		controllerGenerateJsonResponse(
			w,
			r,
			url,
			responses.NewErrorResponse("Unknown error"),
			http.StatusInternalServerError,
		)

		return
	}

	controllerGenerateJsonResponse(
		w,
		r,
		url,
		responses.NewUserSignupResponse(insertedUserDTO),
		http.StatusCreated,
	)
}
