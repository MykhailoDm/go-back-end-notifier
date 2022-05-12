package handler

import (
	"back-end/web/model"
	"back-end/web/service"
)

var us *service.UserService
var ns *service.NotificationService

func LoadServices(m model.Models) {
	us = service.GetUserService(m)
	ns = service.GetNotificationService(m)
}