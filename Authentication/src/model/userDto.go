package model

type UserDto struct {
	Login string `json:"login"`
	Password string `json:"password"`
	AccessToken string `json:"accessToken" param:"token"`
}