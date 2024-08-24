package main

import (
	"encoding/json"
	"fmt"
	"github.com/Nicolas-ggd/go-notification/pkg/storage/models"
	"github.com/nats-io/nats.go"
	"time"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Drain()

	req := []models.Notification{
		{
			Type:    "warning",
			Message: "moicade 2 ",
			Time:    time.Now(),
		},
		{
			Type:    "info",
			Message: "moicade 1",
			Time:    time.Now(),
		},
		{
			Type:    "error",
			Message: "moicade 3",
			Time:    time.Now(),
		},
	}

	val, _ := json.Marshal(&req)

	rep, _ := nc.Request("NOTIFICATION.send-to-all", val, time.Second)
	fmt.Println(string(rep.Data))

}
