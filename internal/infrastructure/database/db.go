package database

import (
	"log"

	"github.com/elonghi/encurtago/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=shorturl port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database ", err)
	}
	db.AutoMigrate(&domain.Link{})
	DB = db

}