package model

type User struct {
	Id string `json:"id"`
	Rating int `json:"rating"`
	TasksSolvedIds []string `json:"tasksSolved"`	
}

type ResponseDto struct {
	Message string `json:"message"`
	User User `json:"user"`
	TopUsers []User `json:"topUsers"`
}

type RequestDto struct {
	UserId string `params:"id" json:"id"`
	RatingPoints int `json:"ratingPoints"`
	TaskId string `json:"solvedTaskId"`
	AccessToken string `params:"token" json:"accessToken"`
}

type AuthResponse struct {
	Message string `json:"message"`
	UserId string `json:"userId"`
	AccessToken string `json:"accessToken"`
}