package main

import (
	"fmt"
	"guestbook_backend/apps/sse_guestbook/routes"
	"guestbook_backend/config"
	natshelper "guestbook_backend/helper/utils/nats_helper"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
)

func main() {

	mux := http.NewServeMux()

	// Connect ke NATS
	nc := config.NewNatsBroker()

	//helper
	natsHelper := natshelper.NewNatsHelper(nc)

	routes.NewRoutes(natsHelper).Routes(mux)

	origins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")
	methods := strings.Split(os.Getenv("CORS_ALLOWED_METHODS"), ",")

	fmt.Println("origins : ", origins)
	fmt.Println("methods : ", methods)

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
