package internal

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		GetEnvVar("DB_HOST"),
		GetEnvVar("DB_USER"),
		GetEnvVar("DB_PASSWORD"),
		GetEnvVar("DB_NAME"),
		GetEnvVar("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}

	DB = db

	db.AutoMigrate(&User{})
}
