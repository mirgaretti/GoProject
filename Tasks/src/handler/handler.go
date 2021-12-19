package handler

import (
	"net/http"
	"strings"
	"tasks/src/model"
	"tasks/src/repository"
	"tasks/src/httpClient"

	"github.com/labstack/echo/v4"
)

func ClassHandler(c echo.Context) error {
	var request model.Request
	c.Bind(request)
	var authResponse model.AuthResponse
	httpClient.GetJson(strings.ReplaceAll("http://localhost:1324?accessToken={accessToken}", "{accessToken}", request.AccessToken), authResponse)
	if authResponse != "valid" {
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