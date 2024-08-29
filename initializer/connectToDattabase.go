package initializer

import (
	"fmt"
	"log"
	"os"
	"simple-api/database"

	"gorm.io/gorm"
)

var PGDB *gorm.DB

func ConnectToDatabase() {
	// Connect to database
	databaseType := os.Getenv("DB_TYPE")
	if databaseType != "postgres" {
		log.Fatalf("Invalid database type: %v", databaseType)
	}
	pgDB, err := database.ConnectToPostgres()
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	fmt.Println("Connected to PostgreSQL")
	PGDB = pgDB

}
