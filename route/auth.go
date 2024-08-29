package route

import (
	"simple-api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine, handler *handler.AuthHandler) {
	// Auth routes
	router.POST("/auth/login", handler.Login)
	router.POST("/auth/logout", handler.Logout)
}
