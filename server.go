package main

import (
	"github.com/labstack/echo/v4"
	"onursume.com/shortener/consts"
	"onursume.com/shortener/router"
)

func main() {
	e := echo.New()
	router.InitializeRouters(e)
	e.Logger.Info(e.Start(":"+ consts.ServerPort))
}