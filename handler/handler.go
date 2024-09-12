package handler

import "simple-api/service"

type HandlerManager struct {
	UserHandler UserHandler
	AuthHandler AuthHandler
	PostHandler PostHandler
}

func NewHandlerManager(services service.ServiceManager) *HandlerManager {
	userHandler := NewUserHandler(services.User)
	authHandler := NewAuthHandler(services.Auth)
	postHandler := NewPostHandler(services.Post)
	return &HandlerManager{
		UserHandler: *userHandler,
		AuthHandler: *authHandler,
		PostHandler: *postHandler,
	}
}
