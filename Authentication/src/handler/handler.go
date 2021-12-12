package handler

import (
	"authentication/src/encryption"
	"authentication/src/model"
	"authentication/src/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context) error {
	u := new(model.UserDto)
	var err error
	if err = c.Bind(u); err != nil {
		return errorHandler(c, err, http.StatusBadRequest)
	}
	user, err := repository.GetUserByLogin(u.Login)
	if err != nil {
		response := model.ResponseDto{
			UserId: "",
			Message: "The login is invalid",
		}
		return c.JSON(http.StatusOK, response)
	}
	isValid := encryption.CheckPassword(u.Password, user.EncryptedPassword)
	if !isValid {
		response := model.ResponseDto{
			UserId: "",
			Message: "The password is invalid",
		}
		return c.JSON(http.StatusOK, response)
	}
	response := model.ResponseDto{
		UserId: user.Id,
		Message: "Success",
	}
	return c.JSON(http.StatusOK, response)
}

func RegisterHandler(c echo.Context) error {
	u := new(model.UserDto)
	var err error
	if err = c.Bind(u); err != nil {
    	return errorHandler(c, err, http.StatusBadRequest)
	}
	enctyptedPassword, err := encryption.HashPassword(u.Password)
	if err != nil {
		return errorHandler(c, err, http.StatusInternalServerError)
	}
	user := model.User{
		Id: uuid.NewString(),
		Login: u.Login,
		EncryptedPassword: enctyptedPassword,
	}

	err = repository.SaveUser(user)
	if err != nil {
		return errorHandler(c, err, http.StatusInternalServerError)
	}

	response := model.ResponseDto{
		UserId: user.Id,
		Message: "Success",
	}
	return c.JSON(http.StatusOK, response)
}

func errorHandler(c echo.Context, err error, status int) error {
	response := model.ResponseDto{
		UserId: "",
		Message: err.Error(),
	}
	return c.JSON(status, response)
}