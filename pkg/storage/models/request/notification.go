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

type IsViewNotificationRequest struct {
	ID     uint `json:"id"`
	IsView bool `json:"is_view"`
}

type ViewNotificationRequest struct {
	ID     uint `json:"id"`
	IsView int  `json:"is_view"`
}

func (nr *NotificationRequest) ToModel() *models.Notification {
	return &models.Notification{
		Type:    nr.Type,
		Message: nr.Message,
		Time:    nr.Time,
		IsView:  nr.IsView,
	}
}

func (nr *IsViewNotificationRequest) ToModel() *ViewNotificationRequest {
	var v ViewNotificationRequest
	v.ID = nr.ID
	if nr.IsView {
		v.IsView = 1
	} else {
		v.IsView = 0
	}

	return &v
}
