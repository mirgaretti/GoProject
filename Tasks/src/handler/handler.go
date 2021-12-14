package handler

import (
	"net/http"
	"tasks/src/model"
	"tasks/src/repository"

	"github.com/labstack/echo/v4"
)

func ClassHandler(c echo.Context) error {
	apikey := c.Request().Header.Get("X-API-KEY")
	if apikey != "8f9c9f1d-860e-4665-9b16-14b2b462e2b2" {
		response := model.ResponseDto {
			Message: "Unauthorized request",
		}
		return c.JSON(http.StatusOK, response)
	}
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