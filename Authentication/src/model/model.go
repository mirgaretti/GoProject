package model

type User struct {
	Id string
	Login string 
	EncryptedPassword string
}

type ResponseDto struct {
	UserId string `json:"userId"`
	AccessToken string `json:"accessToken"`
	Message string `json:"message"`
}

type UserDto struct {
	Login string `json:"login"`
	Password string `json:"password"`
	AccessToken string `json:"accessToken" param:"token"`
}

type UserSession struct {
	UserId string
	AccessToken string
	StartTime string
	ExpireTime string
}