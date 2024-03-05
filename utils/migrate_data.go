package utils

import (
	"log"

	models "github.com/ThanhTien96/airbnb-api/models"

	"gorm.io/gorm"
)

func MigrateData(db *gorm.DB) error {
	log.Println("Migrating data...")
	err := db.AutoMigrate(
		&models.Movie{},
	)
	if err != nil {
		return err
	}
	return nil
}
