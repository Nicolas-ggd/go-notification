package handlers

import (
	"encoding/json"
	"fmt"
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

func NewMicroHandler(service *services.Service) *MicroHandler {
	return &MicroHandler{
		NotificationService: service.NotificationService,
	}
}

func (mh *MicroHandler) BroadcastNotification(wss *ws.Websocket) micro.HandlerFunc {
	return func(r micro.Request) {
		var m models.Notification

		err := json.Unmarshal(r.Data(), &m)
		if err != nil {
			log.Println(err)
		}

		model, err := mh.NotificationService.Insert(&m)
		if err != nil {
			log.Println(err)
		}

		fmt.Printf("%+v\n", model)

		wss.BroadcastEvent(model)
	}
}

func (mh *MicroHandler) ClientBasedNotification(wss *ws.Websocket) micro.HandlerFunc {
	return func(r micro.Request) {
		var m request.NotificationRequest

		err := json.Unmarshal(r.Data(), &m)
		if err != nil {
			log.Println(err)
		}

		model, err := mh.NotificationService.Insert(m.ToModel())
		if err != nil {
			log.Println(err)
		}

		wss.SendEvent(m.Clients, model)
	}
}
