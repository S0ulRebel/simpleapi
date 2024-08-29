package service

import (
	"log"
	"simple-api/repository/postgres"
)

// ServiceManager is a struct to manage all services
type ServiceManager struct {
	User UserService
	Auth AuthService
}

func NewServiceManager(databaseType string) *ServiceManager {
	var userService UserService
	var authService AuthService

	switch databaseType {
	case "postgres":
		userRepo := &postgres.PostgresUserRepository{}
		userService = UserService{
			UserRepo: userRepo,
		}
		authService = AuthService{
			UserRepo: userRepo,
		}
	default:
		log.Fatalf("Invalid database type: %v", databaseType)
	}

	return &ServiceManager{
		User: userService,
		Auth: authService,
	}
}
