package model

type User struct {
	Id int `json:"id"`
	Username string	`json:"username"`
	Password string `json:"password"`
}

type UserInfo struct {
	Id int `json:"id"`
	Username string	`json:"username"`
}