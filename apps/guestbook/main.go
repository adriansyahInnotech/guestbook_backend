package main

import (
	"context"
	"guestbook_backend/apps/guestbook/routes"
	"guestbook_backend/config"
	"guestbook_backend/db"
	"guestbook_backend/helper"
	"log"
	"os"
	"strings"

	"github.com/nats-io/nats.go"
)

func init() {
	db.InitDB()
	db.InitRedis()
	config.InitTracer()
}

func main() {

	defer func() {
		if err := config.ShutdownTracer(context.Background()); err != nil {
			log.Fatalf("failed to shutdown tracer: %v", err)
		}
	}()

	// config.InitCasbin()

	natsBroker := config.NewNatsBroker()
	helper := helper.NewHelper(natsBroker)

	_, err := helper.Utils.Nats.Subscribe("device.*", func(msg *nats.Msg) {
		subject := msg.Subject
		deviceID := strings.TrimPrefix(subject, "device.")
		payload := string(msg.Data)

		helper.Utils.HubSse.Send(deviceID, payload)
	})
	if err != nil {
		log.Fatalf("Error subscribe: %v", err)
	}

	app := config.LoadConfigApp()

	routes.NewRoutes(helper).Routes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))

}
