package request

import "time"

type NotificationRequest struct {
	Type    string    `json:"type"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	Clients *[]string `json:"clients"`
}
