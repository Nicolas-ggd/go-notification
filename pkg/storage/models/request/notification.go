package request

import (
	"github.com/Nicolas-ggd/go-notification/pkg/storage/models"
	"time"
)

type NotificationRequest struct {
	Type    string    `json:"type"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	IsView  bool      `json:"is_view"`
	Clients []string  `json:"clients"`
}

func (nr *NotificationRequest) ToModel() *models.Notification {
	return &models.Notification{
		Type:    nr.Type,
		Message: nr.Message,
		Time:    nr.Time,
		IsView:  nr.IsView,
	}
}
