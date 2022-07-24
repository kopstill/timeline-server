package router

import (
	"kopever/timeline/internal/consts"
	"kopever/timeline/internal/controller"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	setUserController(e)
}

func setUserController(e *echo.Echo) {
	user := controller.NewUserController()
	e.POST(consts.Login, user.Login)
	e.POST(consts.Logout, user.Logout)
}
