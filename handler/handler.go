package handler

import "simple-api/service"

type HandlerManager struct {
	UserHandler UserHandler
	AuthHandler AuthHandler
}

func NewHandlerManager(services service.ServiceManager) *HandlerManager {
	userHandler := NewUserHandler(services.User)
	authHandler := NewAuthHandler(services.Auth)
	return &HandlerManager{
		UserHandler: *userHandler,
		AuthHandler: *authHandler,
	}
}
