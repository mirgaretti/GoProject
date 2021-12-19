package main

import (
	"authentication/src/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/login", handler.LoginHandler)
	e.GET("/validateToken", handler.ValidateTokenHandler)
	e.POST("/register", handler.RegisterHandler)
	e.Logger.Fatal(e.Start(":1324"))
}
