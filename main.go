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
	initializer.LoadEnvironmentVariables()
	initializer.ConnectToDatabase()
	if os.Getenv("DB_TYPE") == "postgres" {
		initializer.SyncDatabase()
	}
}

func main() {
	router := gin.Default()
	services := service.NewServiceManager(os.Getenv("DB_TYPE"))
	handlers := handler.NewHandlerManager(*services)
	route.RegisterRoutes(router, *handlers, *services)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
