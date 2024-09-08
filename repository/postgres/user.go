package postgres

import (
	"fmt"
	"simple-api/errors"
	"simple-api/initializer"
	"simple-api/model"
)

type PostgresUserRepository struct {
}

func (r *PostgresUserRepository) CreateUser(user model.User) (model.User, *errors.AppError) {
	result := initializer.PGDB.Create(&user)
	if result.Error != nil {
		return user, errors.NewErrorService().InternalServerError(result.Error)
	}
	return user, nil
}

func (r *PostgresUserRepository) GetUsers() ([]model.User, *errors.AppError) {
	var users []model.User
	result := initializer.PGDB.Find(&users)
	if result.Error != nil {
		return users, errors.NewErrorService().InternalServerError(result.Error)
	}
	return users, nil
}

func (r *PostgresUserRepository) GetUserByID(id string) (model.User, *errors.AppError) {
	var user model.User
	result := initializer.PGDB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return user, errors.NewErrorService().NotFound("User")
	}
	return user, nil
}

func (r *PostgresUserRepository) GetUserByEmail(email string) (model.User, *errors.AppError) {
	var user model.User
	result := initializer.PGDB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, errors.NewErrorService().NotFound("User")
	}
	return user, nil
}

func (r *PostgresUserRepository) UpdateUser(id string, updatedUser model.User) (model.User, *errors.AppError) {
	var user model.User
	result := initializer.PGDB.First(&user, "id = ?", id)
	if result.Error != nil {
		return user, errors.NewErrorService().NotFound("User")
	}

	updatedUser.ID = user.ID
	fmt.Println(updatedUser)
	result = initializer.PGDB.Model(&user).Updates(&updatedUser)
	if result.Error != nil {
		return user, errors.NewErrorService().InternalServerError(result.Error)
	}

	return user, nil
}

func (r *PostgresUserRepository) DeleteUser(id string) *errors.AppError {
	result := initializer.PGDB.Where("id = ?", id).Delete(&model.User{})
	if result.Error != nil {
		return errors.NewErrorService().InternalServerError(result.Error)
	}
	return nil
}
