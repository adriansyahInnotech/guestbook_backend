package main

import (
	"context"
	"guestbook_backend/apps/guestbook/routes"
	"guestbook_backend/config"
	"guestbook_backend/db"
	"log"
	"os"
)

func init() {
	db.InitDB()
	config.InitTracer()
}

func main() {

	defer func() {
		if err := config.ShutdownTracer(context.Background()); err != nil {
			log.Fatalf("failed to shutdown tracer: %v", err)
		}
	}()

	// config.InitCasbin()

	app := config.LoadConfigApp()
	routes.NewRoutes().Routes(app)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))

}
