package services

import (
	"github.com/Nicolas-ggd/go-notification/pkg/repository"
	"github.com/Nicolas-ggd/go-notification/pkg/storage/models"
)

type INotificationService interface {
	Insert(model *models.Notification) (*models.Notification, error)
}

type NotificationService struct {
	notificationRepository repository.NotificationRepository
}

func NewNotificationService(r repository.NotificationRepository) *NotificationService {
	return &NotificationService{
		notificationRepository: r,
	}
}

func (ns *NotificationService) Insert(model *models.Notification) (*models.Notification, error) {
	model, err := ns.notificationRepository.Insert(model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
