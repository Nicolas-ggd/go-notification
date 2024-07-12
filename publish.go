package main

import (
	"fmt"
	handlers "github.com/Nicolas-ggd/go-notification/pkg/micro_handlers"
	"github.com/nats-io/nats.go"
	"time"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	//notifications := []request.NotificationRequest{
	//	{Type: "warning", Message: "Warning message 1", Time: time.Now(), IsView: true, Clients: []string{"1"}},
	//	{Type: "info", Message: "Info message 1", Time: time.Now(), IsView: true, Clients: []string{"1"}},
	//	{Type: "error", Message: "Error message 1", Time: time.Now(), IsView: false, Clients: []string{"1"}},
	//}

	//by, _ := json.Marshal(notifications)
	rep, _ := nc.Request(handlers.SubjectNotificationList, nil, time.Second)
	fmt.Println(string(rep.Data))

}
