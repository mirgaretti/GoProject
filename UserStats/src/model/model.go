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
}