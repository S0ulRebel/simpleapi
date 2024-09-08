package initializer

import (
	"log"
	"os"
	"simple-api/database"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var PGDB *gorm.DB
var MGDB *mongo.Client

func ConnectToDatabase() {
	// Connect to database
	databaseType := os.Getenv("DB_TYPE")
	switch databaseType {
	case "postgres":
		pgDB, err := database.ConnectToPostgres()
		if err != nil {
			log.Fatalf("Failed to connect to PostgreSQL: %v", err)
		}
		PGDB = pgDB
	case "mongodb":
		mgDB, err := database.ConnectToMongoDB()
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}
		MGDB = mgDB
	default:
		log.Fatalf("Invalid database type: %v", databaseType)

	}
}
