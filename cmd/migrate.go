package main

import (
	"curfile/database"
	"curfile/database/models"
	"log"
)

// RunMigrations ทำการ migrate ตาราง
func main() {
	database.Connect()

	// ทำการ AutoMigrate
	err := database.DB.AutoMigrate(
		&models.Category{},
		&models.Curriculum{},
		&models.Document{},
		&models.DocumentHistory{},
		&models.File{},
		&models.Role{},
		&models.Subject{},
		&models.UserAuth{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	log.Println("Migration completed successfully")
}
