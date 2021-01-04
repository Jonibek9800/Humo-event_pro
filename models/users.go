package models

type Users struct {
	Id       int64  `json:"id"`
	Status   string `json:"status"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Age      int64  `json:"age"`
	Gender   string `json:"gender"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Remove   bool   `json:"remove"`
}

type SignUp struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Age      int64  `json:"age"`
	Gender   string `json:"gender"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ResponseToken struct {
	Description string `json:"description"`
	Token       string `json:"token"`
	}
type LogIn struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
