package route

import (
	"simple-api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, handlers handler.HandlerManager) {
	// Users routes
	RegisterUserRoutes(router, &handlers.UserHandler)
	RegisterAuthRoutes(router, &handlers.AuthHandler)
}
