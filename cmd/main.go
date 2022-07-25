package main

import (
	"kopever/timeline/internal/logger"
	"kopever/timeline/internal/router"

	"github.com/labstack/echo/v4"
)

func main() {
	logger.InitLogger()

	e := echo.New()
	router.Init(e)

	e.Logger.Fatal(e.Start(":1323"))
}
