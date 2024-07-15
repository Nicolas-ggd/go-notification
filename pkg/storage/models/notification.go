// SPDX-License-Identifier: MIT
// Copyright (c) 2024 TOMIOKA
//
// This file is part of the gonotification project.

package models

import "time"

type NotificationType string

const (
	Error   NotificationType = "error"
	Warning NotificationType = "warning"
	Info    NotificationType = "info"
)

type Notification struct {
	ID      uint      `json:"id"`
	Type    string    `json:"type"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}
