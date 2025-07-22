package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB connects to the local PostgreSQL database
func ConnectDB() {
	var err error

	// Example format for PostgreSQL DSN:
	// "host=localhost user=postgres password=yourpassword dbname=mydb port=5432 sslmode=disable"
	dsn := os.Getenv("DB_DSN")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to PostgreSQL database:", err)
	}

	log.Println("Connected to PostgreSQL database successfully")
}

// GetDB returns the DB instance
func GetDB() *gorm.DB {
	return DB
}
