package database

import (
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var DB *gorm.DB

func TestDatabase(t *testing.T) {
	var err error

	dsn := "root:@tcp(127.0.0.1:3306)/koperasi_api?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to Connect to database...")
	}

	fmt.Println("Connectting to database...")
}
