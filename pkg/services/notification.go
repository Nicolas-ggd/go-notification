package services

import (
	"github.com/Nicolas-ggd/go-notification/pkg/repository"
	"github.com/Nicolas-ggd/go-notification/pkg/storage/models"
	"github.com/Nicolas-ggd/go-notification/pkg/storage/models/request"
	metakit "github.com/Nicolas-ggd/gorm-metakit"
)

type INotificationService interface {
	Insert(model *models.Notification) (*models.Notification, error)
	List(meta *metakit.Metadata) (*[]models.Notification, *metakit.Metadata, error)
	Update(model *request.IsViewNotificationRequest) error
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

func (ns *NotificationService) List(meta *metakit.Metadata) (*[]models.Notification, *metakit.Metadata, error) {
	model, meta, err := ns.notificationRepository.List(meta)
	if err != nil {
		return nil, nil, err
	}

	return model, meta, nil
}

func (ns *NotificationService) Update(model *request.IsViewNotificationRequest) error {
	err := ns.notificationRepository.Update(model.ToModel())
	if err != nil {
		return err
	}

	return nil
}
