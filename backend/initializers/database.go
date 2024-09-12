package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/eventhub/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbName)

	if dsn == "" {
		log.Printf("DATABASE_URL not set.")
	}

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}
}

func MigrateDB() {
	err := DB.AutoMigrate(&models.Event{})
	if err != nil {
		log.Printf("Error migrating database: %v", err)
	}
}
