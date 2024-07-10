package services

import "github.com/Nicolas-ggd/go-notification/pkg/repository"

type Service struct {
	NotificationService INotificationService
}

func NewService(repositories *repository.Repository) *Service {
	return &Service{
		NotificationService: NewNotificationService(repositories.NotificationRepository),
	}
}
