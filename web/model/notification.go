package model

type Notification struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Name string `json:"name"`
	UserId int `json:"userId"`
}