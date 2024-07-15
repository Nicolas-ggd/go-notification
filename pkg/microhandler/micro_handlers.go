// SPDX-License-Identifier: MIT
// Copyright (c) 2024 TOMIOKA
//
// This file is part of the gonotification project.

package microhandler

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
	"log"
)

func MicroServices(nc *nats.Conn, handler *MicroHandler) {
	_, err := micro.AddService(nc, micro.Config{
		Name: StreamName,
		Endpoint: &micro.EndpointConfig{
			Subject:    SubjectBroadcastNotification,
			Handler:    handler.BroadcastNotification(),
			Metadata:   nil,
			QueueGroup: "",
		},
		Version:     SubjectVersion,
		Description: "",
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = micro.AddService(nc, micro.Config{
		Name: StreamName,
		Endpoint: &micro.EndpointConfig{
			Subject:    SubjectClientNotification,
			Handler:    handler.ClientBasedNotification(),
			Metadata:   nil,
			QueueGroup: "",
		},
		Version:     SubjectVersion,
		Description: "",
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = micro.AddService(nc, micro.Config{
		Name: StreamName,
		Endpoint: &micro.EndpointConfig{
			Subject:    SubjectNotificationList,
			Handler:    handler.NotificationList(),
			Metadata:   nil,
			QueueGroup: "",
		},
		Version:     SubjectVersion,
		Description: "",
	})
	if err != nil {
		log.Fatal(err)
	}

}
