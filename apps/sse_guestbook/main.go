package main

import (
	"guestbook_backend/apps/sse_guestbook/routes"
	"guestbook_backend/config"
	"guestbook_backend/helper/utils/hub"
	natshelper "guestbook_backend/helper/utils/nats_helper"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nats-io/nats.go"
	"github.com/rs/cors"
)

func main() {

	mux := http.NewServeMux()

	// Connect ke NATS
	nc := config.NewNatsBroker()

	//helper
	natsHelper := natshelper.NewNatsHelper(nc)
	hubSSe := hub.NewHubSse()

	_, err := natsHelper.Subscribe("device.*", func(msg *nats.Msg) {
		subject := msg.Subject
		deviceID := strings.TrimPrefix(subject, "device.")
		hubSSe.Send(deviceID, string(msg.Data))
	})
	if err != nil {
		log.Fatal("failed to subscribe:", err)
	}

	routes.NewRoutes(hubSSe).Routes(mux)

	origins := strings.Split(os.Getenv("CORS_ALLOW_ORIGINS"), ",")
	methods := strings.Split(os.Getenv("CORS_ALLOW_METHOD"), ",")

	c := cors.New(cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   methods,
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Bungkus mux dengan CORS handler
	handler := c.Handler(mux)

	log.Printf("Server SSE running at http://localhost%s", os.Getenv("APP_PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("APP_PORT"), handler))

}
