package route

import (
	"simple-api/handler"
	"simple-api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, handler *handler.UserHandler) {

	// Users routes
	router.POST("/users", handler.CreateUser)
	router.GET("/users", middleware.RequireAuth, handler.GetUsers)
	router.GET("/users/:id", middleware.RequireAuth, handler.GetUserByID)
	router.PUT("/users/:id", middleware.RequireAuth, handler.UpdateUser)
	router.DELETE("/users/:id", middleware.RequireAuth, handler.DeleteUser)
}
