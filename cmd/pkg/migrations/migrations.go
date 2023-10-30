package migrations

import (
	"log"

	"github.com/Taylorgtyler/go-rest-api/cmd/pkg/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Activity{},
		&models.List{},
		&models.User{})

	if err != nil {
		log.Fatalf("Could not migrate models: %v", err)
	}
	log.Println("Migrations ran successfully")
}
