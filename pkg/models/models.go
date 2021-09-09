package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SignUp struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
