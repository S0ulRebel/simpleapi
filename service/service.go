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
	Post PostService
}

func NewServiceManager(databaseType string) *ServiceManager {
	var userRepo repository.UserRepository
	var postRepo repository.PostRepository
	var userService UserService
	var authService AuthService
	var postService PostService

	switch databaseType {
	case "postgres":
		userRepo = &postgres.PostgresUserRepository{}
		postRepo = &postgres.PostgresPostRepository{}
	case "mongodb":
		userRepo = &mongo.MongoUserRepository{}
		postRepo = &mongo.MongoPostRepository{}
	default:
		log.Fatalf("Invalid database type: %v", databaseType)
	}

	userService = UserService{
		UserRepo: userRepo,
	}
	authService = AuthService{
		UserRepo: userRepo,
	}
	postService = PostService{
		PostRepo: postRepo,
	}

	return &ServiceManager{
		User: userService,
		Auth: authService,
		Post: postService,
	}
}
