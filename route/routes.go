package route

import (
	"simple-api/handler"
	"simple-api/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, handlers handler.HandlerManager, services service.ServiceManager) {
	RegisterUserRoutes(router, &handlers.UserHandler, services.User)
	RegisterAuthRoutes(router, &handlers.AuthHandler)
	RegisterPostRoutes(router, &handlers.PostHandler, services.User)
}
