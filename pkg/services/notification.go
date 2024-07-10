package services

import "github.com/Nicolas-ggd/go-notification/pkg/repository"

type INotificationService interface{}

type NotificationService struct {
	notificationRepository repository.NotificationRepository
}

func NewNotificationService(r repository.NotificationRepository) *NotificationService {
	return &NotificationService{
		notificationRepository: r,
	}
}
