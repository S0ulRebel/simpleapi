package service

import (
	"simple-api/errors"
	"simple-api/model"
	"simple-api/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestUserService_CreateUser_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userService := UserService{UserRepo: mockRepo}

	inputUser := model.User{
		ID:       "1",
		Email:    "testuser@test.com",
		Password: "password123",
	}

	mockRepo.On("CreateUser", mock.Anything).Return(inputUser, nil)

	createdUser, appErr := userService.CreateUser(inputUser)

	if appErr != nil {
		t.Errorf("expected no error, got %v", appErr)
	}
	if createdUser.Email != inputUser.Email {
		t.Errorf("expected email %s, got %s", inputUser.Email, createdUser.Email)
	}
	if createdUser.ID == "" {
		t.Errorf("expected non-empty ID, got %s", createdUser.ID)
	}

	mockRepo.AssertExpectations(t)
}

func TestUserService_CreateUser_PasswordHashError(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userService := UserService{UserRepo: mockRepo}

	inputUser := model.User{
		Email:    "testuser@test.com",
		Password: string(make([]byte, bcrypt.MaxCost+1)),
	}

	mockRepo.On("CreateUser", mock.Anything).Return(model.User{}, nil)

	createdUser, appErr := userService.CreateUser(inputUser)

	assert.Error(t, appErr)
	assert.Equal(t, "", createdUser.ID)
	mockRepo.AssertNotCalled(t, "CreateUser")
}

func TestUserService_CreateUser_RepoError(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userService := UserService{UserRepo: mockRepo}

	inputUser := model.User{
		Email:    "testuser@test.com",
		Password: "password123",
	}

	mockRepo.On("CreateUser", mock.Anything).Return(model.User{}, errors.NewErrorService().InternalServerError(
		&errors.AppError{
			Code:    500,
			Message: "error creating user",
			Err:     nil,
		},
	))
	createdUser, appErr := userService.CreateUser(inputUser)

	// Assert
	assert.Error(t, appErr)
	assert.Equal(t, "", createdUser.ID)
	mockRepo.AssertExpectations(t)
}
