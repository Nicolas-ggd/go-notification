// SPDX-License-Identifier: MIT
// Copyright (c) 2024 TOMIOKA
//
// This file is part of the go-notification project.

package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Nicolas-ggd/go-notification/pkg/http/ws"
	"github.com/Nicolas-ggd/go-notification/pkg/queue"
	"github.com/Nicolas-ggd/go-notification/pkg/services"
	"github.com/Nicolas-ggd/go-notification/pkg/storage/models"
	"github.com/Nicolas-ggd/go-notification/pkg/storage/models/request"
	metakit "github.com/Nicolas-ggd/gorm-metakit"
	"github.com/nats-io/nats.go/micro"
	"log"
)

type MicroHandler struct {
	NotificationService services.INotificationService
	PriorityQueue       *queue.PriorityQueue
	wss                 *ws.Websocket
}

func NewMicroHandler(service *services.Service, pr *queue.PriorityQueue, wss *ws.Websocket) *MicroHandler {
	return &MicroHandler{
		NotificationService: service.NotificationService,
		PriorityQueue:       pr,
		wss:                 wss,
	}
}

func (mh *MicroHandler) BroadcastNotification() micro.HandlerFunc {
	return func(r micro.Request) {
		var m []models.Notification

		err := json.Unmarshal(r.Data(), &m)
		if err != nil {
			log.Println(err)
		}

		for _, notification := range m {
			model, err := mh.NotificationService.Insert(&notification)
			if err != nil {
				log.Println(err)
			}

			// Push the go-notification into the priority queue
			mh.PriorityQueue.Push(&queue.NotificationHeap{
				ID:      model.ID,
				Type:    model.Type,
				Message: model.Message,
				Time:    model.Time,
				IsView:  model.IsView,
			})
		}

		// process queue and send broadcast event
		mh.processQueue()
	}
}

func (mh *MicroHandler) ClientBasedNotification() micro.HandlerFunc {
	return func(r micro.Request) {
		var m []request.NotificationRequest

		err := json.Unmarshal(r.Data(), &m)
		if err != nil {
			log.Println(err)
		}

		for _, notification := range m {
			model, err := mh.NotificationService.Insert(notification.ToModel())
			if err != nil {
				log.Println(err)
			}

			// Push the go-notification into the priority queue
			mh.PriorityQueue.Push(&queue.NotificationHeap{
				ID:      model.ID,
				Type:    model.Type,
				Message: model.Message,
				Time:    model.Time,
				IsView:  model.IsView,
				Clients: notification.Clients,
			})
		}

		// process queue and send client based event
		mh.processQueue()
	}
}

func (mh *MicroHandler) NotificationList() micro.HandlerFunc {
	metadata := &metakit.Metadata{Sort: "id"}
	return func(r micro.Request) {
		model, _, err := mh.NotificationService.List(metadata)
		if err != nil {
			log.Println(err)
		}

		fmt.Printf("%+v\n", model)
	}
}

func (mh *MicroHandler) NotificationViewed() micro.HandlerFunc {
	return func(r micro.Request) {
		var m request.IsViewNotificationRequest

		err := json.Unmarshal(r.Data(), &m)
		if err != nil {
			log.Println(err)
		}

		err = mh.NotificationService.Update(&m)
		if err != nil {
			log.Println(err)
		}
	}
}

func (mh *MicroHandler) processQueue() {
	for mh.PriorityQueue.Len() > 0 {
		notification := mh.PriorityQueue.Pop().(*queue.NotificationHeap)
		model := &models.Notification{
			ID:      notification.ID,
			Type:    notification.Type,
			Message: notification.Message,
			Time:    notification.Time,
			IsView:  notification.IsView,
		}

		// check if there are specific clients to send the go-notification to
		if len(notification.Clients) > 0 {
			// send event to specific clients
			mh.wss.SendEvent(notification.Clients, model)
		} else {
			// Broadcast event to all clients
			mh.wss.BroadcastEvent(model)
		}
	}
}
