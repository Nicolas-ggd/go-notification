package handlers

import (
	"encoding/json"
	"github.com/Nicolas-ggd/go-notification/pkg/http/ws"
	"github.com/Nicolas-ggd/go-notification/pkg/services"
	"github.com/Nicolas-ggd/go-notification/pkg/storage/models"
	"github.com/Nicolas-ggd/go-notification/pkg/storage/models/request"
	"github.com/nats-io/nats.go/micro"
	"log"
)

type MicroHandler struct {
	NotificationService services.INotificationService
}

func NewMicroHandler(notificationService services.INotificationService) *MicroHandler {
	return &MicroHandler{
		NotificationService: notificationService,
	}
}

func BroadcastNotification(wss *ws.Websocket) micro.HandlerFunc {
	return func(r micro.Request) {
		var m models.Notification

		err := json.Unmarshal(r.Data(), &m)
		if err != nil {
			log.Println(err)
		}

		wss.BroadcastEvent(r.Data())
	}
}

func ClientBasedNotification(wss *ws.Websocket) micro.HandlerFunc {
	return func(r micro.Request) {
		var m request.NotificationRequest

		err := json.Unmarshal(r.Data(), &m)
		if err != nil {
			log.Println(err)
		}

		wss.SendEvent(m.Clients, r.Data())
	}
}
