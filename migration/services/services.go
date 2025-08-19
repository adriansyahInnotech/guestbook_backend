package services

import (
	registrationmodel "guestbook/migration/registration_model"
	"log"

	"gorm.io/gorm"
)

func MigrateUp(db *gorm.DB) {
	for _, model := range registrationmodel.RegisterModels() {
		err := db.AutoMigrate(model)
		if err != nil {
			log.Fatalf("❌ Failed to migrate model %T: %v", model, err)
		}
	}
	log.Println("✅ All models migrated successfully (UP)")
}

func MigrateDown(db *gorm.DB) {
	for _, model := range registrationmodel.RegisterModels() {
		err := db.Migrator().DropTable(model)
		if err != nil {
			log.Fatalf("❌ Failed to drop table for model %T: %v", model, err)
		}
	}
	log.Println("🧨 All tables dropped successfully (DOWN)")
}
