package service

import (
	"simple-api/model"
	"simple-api/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func (s *UserService) CreateUser(user model.User) (model.User, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password),10)
	if err != nil {
		return model.User{}, err
	}
	user.Password = string(hashedPwd)
	return s.UserRepo.CreateUser(user)
}

func (s *UserService) GetUsers() ([]model.User, error) {
	return s.UserRepo.GetUsers()
}

func (s *UserService) GetUserByID(id int) (model.User, error) {
	return s.UserRepo.GetUserByID(id)
}

func (s *UserService) UpdateUser(id int, updatedUser model.User) (model.User, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password),10)
	if err != nil {
		return model.User{}, err
	}
	updatedUser.Password = string(hashedPwd)
	return s.UserRepo.UpdateUser(id, updatedUser)
}

func (s *UserService) DeleteUser(id int) error {
	return s.UserRepo.DeleteUser(id)
}
