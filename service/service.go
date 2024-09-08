package service

import (
	"log"
	"simple-api/repository"
	"simple-api/repository/mongo"
	"simple-api/repository/postgres"
)

// ServiceManager is a struct to manage all services
type ServiceManager struct {
	User UserService
	Auth AuthService
}

func NewServiceManager(databaseType string) *ServiceManager {
	var userRepo repository.UserRepository
	var userService UserService
	var authService AuthService

	switch databaseType {
	case "postgres":
		userRepo = &postgres.PostgresUserRepository{}
	case "mongodb":
		userRepo = &mongo.MongoUserRepository{}
	default:
		log.Fatalf("Invalid database type: %v", databaseType)
	}

	userService = UserService{
		UserRepo: userRepo,
	}
	authService = AuthService{
		UserRepo: userRepo,
	}

	return &ServiceManager{
		User: userService,
		Auth: authService,
	}
}
