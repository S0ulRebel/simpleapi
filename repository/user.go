package repository

import (
	"simple-api/model"
)

type UserRepository interface {
	CreateUser(user model.User) (model.User, error)
	GetUsers() ([]model.User, error)
	GetUserByID(id int) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
	UpdateUser(id int, updatedUser model.User) (model.User, error)
	DeleteUser(id int) error
}
