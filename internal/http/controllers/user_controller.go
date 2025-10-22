package controllers

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/dto"
	requests "demo-todo-manager/internal/http/requests/user"
	responses "demo-todo-manager/internal/http/responses/user"
	"net/http"
)

type userController struct {
	userService contracts.UserService
	authService contracts.AuthService
}

func newUserController(userService contracts.UserService, authService contracts.AuthService) contracts.UserController {
	return &userController{
		userService: userService,
		authService: authService,
	}
}

// @Summary User login
// @Description Logs in a user and returns a token
// @Tags auth
// @Accept json
// @Produce json
// @Param Accept-Language header string false "User locale" Enums(ru, en) Example(en) Default(en)
// @Param request body requests.UserLoginRequest true "Login email and password"
// @Success 200 {object} responses.UserLoginResponseStruct
// @Failure 400 {object} responses.ErrorResponseStruct
// @Failure 401 {object} responses.ErrorResponseStruct
// @Failure 422 {object} responses.ValidationErrorResponseStruct
// @Failure 500 {object} responses.ErrorResponseStruct
// @Router /login [post]
func (c *userController) Login(w http.ResponseWriter, r *http.Request) {
	var req requests.UserLoginRequest

	if !Preparation(w, r, &req, requests.UserLoginValidateMethod) {
		return
	}

	userDTO := dto.UserDTO{
		Email:    req.Email,
		Password: req.Password,
	}

	existedUserDTO, ok := c.userService.GetByEmail(userDTO.Email)
	if !ok {
		UnknownErrorResponse(w, r)
		return
	}

	localizer := GetLocalizer(r)

	if (existedUserDTO == dto.UserDTO{}) {
		UnauthorizedResponse(w, r, localizer.T("auth.login_failed", nil))
		return
	}

	if !c.userService.ValidatePassword(userDTO.Password, existedUserDTO.Password) {
		UnauthorizedResponse(w, r, localizer.T("auth.login_failed", nil))
		return
	}

	token, err := c.authService.IssueToken(existedUserDTO.ID, true, true)
	if err != nil {
		UnknownErrorResponse(w, r)
		return
	}

	JsonResponse(
		w,
		r,
		responses.UserLoginResponse(token, localizer.T("auth.login_succeded", map[string]interface{}{
			"email": userDTO.Email,
		})),
		http.StatusOK,
	)
}

// @Summary User signup
// @Description Registers a user and returns it's data
// @Tags auth
// @Accept json
// @Produce json
// @Param Accept-Language header string false "User locale" Enums(ru, en) Example(en) Default(en)
// @Param request body requests.UserSignupRequest true "Signup email and password"
// @Success 200 {object} responses.UserSignupResponseStruct
// @Failure 400 {object} responses.ErrorResponseStruct
// @Failure 401 {object} responses.ErrorResponseStruct
// @Failure 409 {object} responses.ErrorResponseStruct
// @Failure 422 {object} responses.ValidationErrorResponseStruct
// @Failure 500 {object} responses.ErrorResponseStruct
// @Router /login [post]
func (c *userController) Signup(w http.ResponseWriter, r *http.Request) {
	var req requests.UserSignupRequest

	if !Preparation(w, r, &req, requests.UserSignupValidateMethod) {
		return
	}

	userDTO := dto.UserDTO{
		Email:    req.Email,
		Password: req.Password,
	}

	existedUserDTO, ok := c.userService.GetByEmail(userDTO.Email)
	if !ok {
		UnknownErrorResponse(w, r)
		return
	}

	localizer := GetLocalizer(r)

	if (existedUserDTO != dto.UserDTO{}) {
		ConflictResponse(w, r, localizer.T("auth.user_exists", map[string]interface{}{
			"email": userDTO.Email,
		}))
		return
	}

	insertedUserDTO, err := c.userService.Store(userDTO)
	if err != nil {
		UnknownErrorResponse(w, r)
		return
	}

	JsonResponse(
		w,
		r,
		responses.UserSignupResponse(insertedUserDTO, localizer.T("user.created_successfully", map[string]interface{}{
			"email": userDTO.Email,
		})),
		http.StatusCreated,
	)
}
