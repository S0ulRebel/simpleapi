package route

import (
	"simple-api/handler"
	"simple-api/middleware"
	"simple-api/service"

	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(
	router *gin.Engine,
	handler *handler.PostHandler,
	userService service.UserService,
) {
	// Post routes
	router.POST("/posts", middleware.RequireAuth(userService), handler.CreatePost)
	router.GET("/posts", handler.GetPosts)
	router.GET("/posts/:id", handler.GetPostByID)
	router.GET("/posts/user/:userID", handler.GetPostsByUserID)
	router.PUT("/posts/:id", middleware.RequireAuth(userService), handler.UpdatePost)
	router.DELETE("/posts/:id", middleware.RequireAuth(userService), handler.DeletePost)
}
