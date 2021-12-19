package model

type ResponseDto struct {
	UserId string `json:"userId"`
	AccessToken string `json:"accessToken"`
	Message string `json:"message"`
}