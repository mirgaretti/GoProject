package handler

import (
	"authentication/src/encryption"
	"authentication/src/model"
	"authentication/src/repository"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const layout string = "2006-01-02T15:04:05.000Z"

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
			AccessToken: "",
			Message: "The login is invalid",
		}
		return c.JSON(http.StatusOK, response)
	}
	isValid := encryption.CheckPassword(u.Password, user.EncryptedPassword)
	if !isValid {
		response := model.ResponseDto{
			UserId: "",
			AccessToken: "",
			Message: "The password is invalid",
		}
		return c.JSON(http.StatusOK, response)
	}

	currentTime := time.Now();
	session := model.UserSession {
		UserId: user.Id,
		AccessToken: uuid.NewString(),
		StartTime: currentTime.Format(layout),
		ExpireTime: currentTime.Local().Add(time.Hour * time.Duration(1)).Format(layout),
	}
	err = repository.SaveSession(session)
	if err != nil {
		return err
	}
	response := model.ResponseDto{
		UserId: user.Id,
		AccessToken: "",
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
		AccessToken: "",
		Message: "Success",
	}
	return c.JSON(http.StatusOK, response)
}

func ValidateTokenHandler(c echo.Context) error {
	u := new(model.UserDto)
	var err error
	if err = c.Bind(u); err != nil {
    	return errorHandler(c, err, http.StatusBadRequest)
	}
	session, err := repository.GetSessionByToken(u.AccessToken)
	if err != nil {
		response := model.ResponseDto{
			UserId: "",
			AccessToken: "",
			Message: "invalid",
		}
		return c.JSON(http.StatusOK, response)
	}

	expireTime, err := time.Parse(layout, session.ExpireTime)
	if err != nil {
		response := model.ResponseDto{
			UserId: "",
			AccessToken: "",
			Message: "invalid",
		}
		return c.JSON(http.StatusOK, response)
	}

	isValid := expireTime.After(time.Now())

	if isValid {
		response := model.ResponseDto{
			UserId: "",
			AccessToken: "",
			Message: "valid",
		}
		return c.JSON(http.StatusOK, response)
	}

	response := model.ResponseDto{
		UserId: "",
		AccessToken: "",
		Message: "invalid",
	}
	return c.JSON(http.StatusOK, response)
}

func errorHandler(c echo.Context, err error, status int) error {
	response := model.ResponseDto{
		UserId: "",
		AccessToken: "", 
		Message: err.Error(),
	}
	return c.JSON(status, response)
}