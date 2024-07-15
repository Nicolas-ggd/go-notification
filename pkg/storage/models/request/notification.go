// SPDX-License-Identifier: MIT
// Copyright (c) 2024 TOMIOKA
//
// This file is part of the gonotification project.

package request

import (
	"github.com/Nicolas-ggd/go-notification/pkg/storage/models"
	"time"
)

type NotificationRequest struct {
	ID      uint      `json:"id"`
	Type    string    `json:"type"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	Clients []string  `json:"clients"`
}

func (nr *NotificationRequest) ToModel() *models.Notification {
	return &models.Notification{
		Type:    nr.Type,
		Message: nr.Message,
		Time:    nr.Time,
	}
}
