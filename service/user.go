package service

import (
	"simple-api/errors"
	"simple-api/model"
	"simple-api/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func (s *UserService) CreateUser(user model.User) (model.User, *errors.AppError) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.User{}, errors.NewErrorService().InternalServerError(err)
	}
	user.Password = string(hashedPwd)
	return s.UserRepo.CreateUser(user)
}

func (s *UserService) GetUsers() ([]model.User, *errors.AppError) {
	return s.UserRepo.GetUsers()
}

func (s *UserService) GetUserByID(id int) (model.User, *errors.AppError) {
	return s.UserRepo.GetUserByID(id)
}

func (s *UserService) UpdateUser(id int, updatedUser model.User) (model.User, *errors.AppError) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), 10)
	if err != nil {
		return model.User{}, errors.NewErrorService().InternalServerError(err)
	}
	updatedUser.Password = string(hashedPwd)
	return s.UserRepo.UpdateUser(id, updatedUser)
}

func (s *UserService) DeleteUser(id int) *errors.AppError {
	return s.UserRepo.DeleteUser(id)
}
