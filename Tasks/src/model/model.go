package model

type Task struct {
	Id string `json:"id"`
	Label string `json:"label"`
	TopicLabel string
	Text string `json:"text"`
	Answer string `json:"answer"`
	Rating int `json:"rating"`
}

type Topic struct {
	Label string `json:"label"`
	ClassLabel int
	Description string `json:"description"`
	Tasks []Task `json:"tasks"`
}

type Class struct {
	Label int16 `json:"label"`
	Description string `json:"description"`
	Topics []Topic `json:"topics"`
}

type ResponseDto struct {
	Message string `json:"message"`
	Classes []Class `json:"classes"`
}