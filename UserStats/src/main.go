package main

import (
	"userStats/src/handler"

	"github.com/labstack/echo/v4"
)

func main() {	
	e := echo.New()
	e.GET("/user", handler.GetUserHandler)
	e.GET("/top", handler.GetTopHandler)
	e.POST("/user", handler.TaskUserHandler)
	e.POST("/user/add", handler.AddUserHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
