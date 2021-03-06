package handler

import (
	"strings"
	"net/http"
	"userStats/src/model"
	"userStats/src/repository"
	"userStats/src/httpClient"

	"github.com/labstack/echo/v4"
)

func GetUserHandler(c echo.Context) error {
	var authResponse model.AuthResponse
	httpClient.GetJson(strings.ReplaceAll("http://localhost:1324?accessToken={accessToken}", "{accessToken}", request.AccessToken), authResponse)
	if authResponse != "valid" {
		response := model.ResponseDto {
			Message: "Unauthorized request",
		}
		return c.JSON(http.StatusOK, response)
	}
	var request model.RequestDto
	err := c.Bind(request);
	if err != nil {
		return err
	}
	user, err := repository.GetUserById(request.UserId)

	response := model.ResponseDto {
		Message: "Success",
		User: user,
	}
	return c.JSON(http.StatusOK, response)
}

func GetTopHandler(c echo.Context) error {
	var request model.RequestDto
	err := c.Bind(request);
	if err != nil {
		return err
	}
	users, err := repository.GetTopUserIds()

	response := model.ResponseDto {
		Message: "Success",
		TopUsers: users,
	}
	return c.JSON(http.StatusOK, response)
}

func TaskUserHandler(c echo.Context) error {
	var authResponse model.AuthResponse
	httpClient.GetJson(strings.ReplaceAll("http://localhost:1324?accessToken={accessToken}", "{accessToken}", request.AccessToken), authResponse)
	if authResponse != "valid" {
		response := model.ResponseDto {
			Message: "Unauthorized request",
		}
		return c.JSON(http.StatusOK, response)
	}
	var request model.RequestDto
	err := c.Bind(request);
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	user, err := repository.GetUserById(request.UserId)
	user.Rating = user.Rating + request.RatingPoints
	err = repository.SaveUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	response := model.ResponseDto {
		Message: "Success",
		User: user,
	}
	return c.JSON(http.StatusOK, response)
}

func AddUserHandler(c echo.Context) error {
	var authResponse model.AuthResponse
	httpClient.GetJson(strings.ReplaceAll("http://localhost:1324?accessToken={accessToken}", "{accessToken}", request.AccessToken), authResponse)
	if authResponse != "valid" {
		response := model.ResponseDto {
			Message: "Unauthorized request",
		}
		return c.JSON(http.StatusOK, response)
	}
	var request model.RequestDto
	err := c.Bind(request);
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	user := model.User {
		Id: request.UserId,
		Rating: 0,
	}
	err = repository.SaveUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	response := model.ResponseDto {
		Message: "Success",
		User: user,
	}
	return c.JSON(http.StatusOK, response)
}