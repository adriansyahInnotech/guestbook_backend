package main

import (
	"flag"
	"fmt"
	"guestbook_backend/db"
	"guestbook_backend/migration/services"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()

}

func main() {
	var action string
	flag.StringVar(&action, "action", "up", "Migration action: up or down")
	flag.Parse()

	fmt.Println("database_url : ", os.Getenv("DATABASE_URL"))

	switch action {
	case "up":
		services.MigrateUp(db.GetDB())
	case "down":
		services.MigrateDown(db.GetDB())
	default:
		log.Println("❓ Invalid action. Use -action=up or -action=down")
	}
}
