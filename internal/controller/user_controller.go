package controller

import (
	"kopever/timeline/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	Login(echo.Context) error
	Logout(echo.Context) error
}

type userController struct {
	service service.UserService
}

func NewUserController() UserController {
	return &userController{service.NewUserService()}
}

func (controller *userController) Login(e echo.Context) error {
	e.String(http.StatusOK, "login")
	return controller.service.Login()
}

func (controller *userController) Logout(e echo.Context) error {
	e.String(http.StatusOK, "logout")
	return controller.service.Logout()
}
