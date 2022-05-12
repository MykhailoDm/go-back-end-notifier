package service

import (
	"back-end/web/model"
	"log"
	"sync"
)

type NotificationService struct {
	models model.Models
}

var notificationLock = &sync.Mutex{}
var ns *NotificationService

func GetNotificationService(m model.Models) *NotificationService {
	if ns == nil {
		notificationLock.Lock()
		defer notificationLock.Unlock()
		if ns == nil {
			log.Println("Creating user service instance")
			ns = &NotificationService{
				models: m,
			}
		}
	}

	return ns
}

func (n *NotificationService) GetNotifications(uid int) ([]*model.Notification, error) {
	log.Printf("Retrieving notifications for user: %v", uid)
	return n.models.DB.GetNotifications(uid)
}
