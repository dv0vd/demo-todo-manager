package utils

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/http/controllers"
)

func ControllerInitControllers(userService contracts.UserService, authService contracts.AuthService) (contracts.UserController, contracts.AuthController) {
	return controllers.NewUserController(userService, authService), controllers.NewAuthController(authService)
}
