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

func (ns *NotificationService) GetNotifications(uid int) ([]*model.Notification, error) {
	log.Printf("Retrieving notifications for user: %v", uid)
	return ns.models.DB.GetNotifications(uid)
}

func (ns *NotificationService) GetNotification(id int) (*model.Notification, error) {
	log.Printf("Retrieving notification with id: %v", id)
	return ns.models.DB.GetNotification(id)
}

func (ns *NotificationService) CreateNotification(n model.Notification) error {
	return ns.models.DB.CreateNotification(n)
}

func (ns *NotificationService) DeleteNotification(id int) error {
	return ns.models.DB.DeleteNotification(id)
}

func (ns *NotificationService) UpdateNotification(id int, n model.Notification) error {
	return ns.models.DB.UpdateNotification(id, n)
}
