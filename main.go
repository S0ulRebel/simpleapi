package main

import (
	"fmt"
	"os"
	"simple-api/handler"
	"simple-api/initializer"
	"simple-api/route"
	"simple-api/service"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Starting the application...")
	initializer.LoadEnvironmentVariables()
	initializer.ConnectToDatabase()
	initializer.SyncDatabase()
}

func main() {
	// Initialize gin router
	router := gin.Default()
	// Initialize services
	services := service.NewServiceManager(os.Getenv("DB_TYPE"))
	// Initialize handlers
	handlers := handler.NewHandlerManager(*services)
	// Register routes
	route.RegisterRoutes(router, *handlers)
	router.Run(":8080")
}
