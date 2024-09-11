package mocks

import (
	"simple-api/errors"
	"simple-api/model"

	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (m *UserRepository) CreateUser(user model.User) (model.User, *errors.AppError) {
	args := m.Called(user)
	userResult := args.Get(0).(model.User)
	appErrResult := args.Get(1)
	if appErrResult == nil {
		return userResult, nil
	}
	return userResult, appErrResult.(*errors.AppError)
}

func (m *UserRepository) GetUsers() ([]model.User, *errors.AppError) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Get(1).(*errors.AppError)
}

func (m *UserRepository) GetUserByID(id string) (model.User, *errors.AppError) {
	args := m.Called(id)
	return args.Get(0).(model.User), args.Get(1).(*errors.AppError)
}

func (m *UserRepository) GetUserByEmail(email string) (model.User, *errors.AppError) {
	args := m.Called(email)
	return args.Get(0).(model.User), args.Get(1).(*errors.AppError)
}

func (m *UserRepository) UpdateUser(id string, updatedUser model.User) (model.User, *errors.AppError) {
	args := m.Called(id, updatedUser)
	return args.Get(0).(model.User), args.Get(1).(*errors.AppError)
}

func (m *UserRepository) DeleteUser(id string) *errors.AppError {
	args := m.Called(id)
	return args.Get(0).(*errors.AppError)
}
