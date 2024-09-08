package repository

import (
	"simple-api/errors"
	"simple-api/model"
)

type UserRepository interface {
	CreateUser(user model.User) (model.User, *errors.AppError)
	GetUsers() ([]model.User, *errors.AppError)
	GetUserByID(id string) (model.User, *errors.AppError)
	GetUserByEmail(email string) (model.User, *errors.AppError)
	UpdateUser(id string, updatedUser model.User) (model.User, *errors.AppError)
	DeleteUser(id string) *errors.AppError
}
