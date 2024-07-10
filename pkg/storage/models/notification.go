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
	IsView  bool      `json:"is_view"`
}
