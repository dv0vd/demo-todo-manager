package utils

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/http/controllers"
)

func ControllerInitControllers(userService contracts.UserService) contracts.UserController {
	return controllers.NewUserController(userService)
}
