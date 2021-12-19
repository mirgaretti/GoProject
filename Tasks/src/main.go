package main

import (
	"tasks/src/handler"

	"github.com/labstack/echo/v4"
)

func main() {	
	e := echo.New()
	e.GET("/class", handler.ClassHandler)
	e.Logger.Fatal(e.Start(":1325"))
}
