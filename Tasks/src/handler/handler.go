package handler

import (
	"net/http"
	"tasks/src/model"
	"tasks/src/repository"

	"github.com/labstack/echo/v4"
)

func ClassHandler(c echo.Context) error {
	
	classes, err := repository.GetAllClasses()
	if err != nil{
		response := model.ResponseDto {
			Message: err.Error(),
		}
		return c.JSON(http.StatusOK, response)
	}
	response := model.ResponseDto {
		Message: "Success",
		Classes: classes,
	}
	return c.JSON(http.StatusOK, response)
}