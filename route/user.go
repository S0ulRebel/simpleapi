package route

import (
	"simple-api/handler"
	"simple-api/middleware"
	"simple-api/service"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, handler *handler.UserHandler, userService service.UserService) {

	// Users routes
	router.POST("/users", handler.CreateUser)
	router.GET("/users", middleware.RequireAuth(userService), handler.GetUsers)
	router.GET("/users/:id", middleware.RequireAuth(userService), handler.GetUserByID)
	router.PUT("/users/:id", middleware.RequireAuth(userService), handler.UpdateUser)
	router.DELETE("/users/:id", middleware.RequireAuth(userService), handler.DeleteUser)
}
