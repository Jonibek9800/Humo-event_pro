package models

type Vacancy struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Salary      string `json:"salary"`
	Description string `json:"description"`
	DataAdd     string `json:"data_add"`
}
