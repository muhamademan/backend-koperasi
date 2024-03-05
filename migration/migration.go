package migration

import (
	"backend-koperasi/database"
	"backend-koperasi/models/entity"
	"fmt"
	"log"
)

func Migrate() {
	err := database.DB.AutoMigrate(
		&entity.User{},
	)

	// jika error
	if err != nil {
		log.Fatal("Failed to migrate...")
	}
	// Jika success
	fmt.Println("Migration successful")
}
