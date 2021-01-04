package models

type News struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Data     string `json:"data"`
	Textarea string `json:"textarea"`
}
