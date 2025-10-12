package utils

import (
	"demo-todo-manager/internal/contracts"
	"demo-todo-manager/internal/http/controllers"
)

func ControllerInitControllers(userService contracts.UserService, authService contracts.AuthService, noteService contracts.NoteService) (contracts.UserController, contracts.AuthController, contracts.NoteController) {
	return controllers.NewUserController(userService, authService), controllers.NewAuthController(authService), controllers.NewNoteController(authService, userService, noteService)
}
