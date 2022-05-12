package handler

import (
	"back-end/web/model"
	"back-end/web/service"
)

var us *service.UserService

func LoadServices(m model.Models) {
	us = service.GetUserService(m)
}