package handler

import (
	"net/http"
	"userStats/src/model"
	"userStats/src/repository"

	"github.com/labstack/echo/v4"
)

func GetUserHandler(c echo.Context) error {
	apikey := c.Request().Header.Get("X-API-KEY")
	if apikey != "8f9c9f1d-860e-4665-9b16-14b2b462e2b2" {
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
	apikey := c.Request().Header.Get("X-API-KEY")
	if apikey != "8f9c9f1d-860e-4665-9b16-14b2b462e2b2" {
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
	apikey := c.Request().Header.Get("X-API-KEY")
	if apikey != "8f9c9f1d-860e-4665-9b16-14b2b462e2b2" {
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