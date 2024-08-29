package postgres

import (
	"simple-api/initializer"
	"simple-api/model"
)

type PostgresUserRepository struct {
}

func (r *PostgresUserRepository) CreateUser(user model.User) (model.User, error) {
	result := initializer.PGDB.Create(&user)
	return user, result.Error
}

func (r *PostgresUserRepository) GetUsers() ([]model.User, error) {
	var users []model.User
	result := initializer.PGDB.Find(&users)
	return users, result.Error
}

func (r *PostgresUserRepository) GetUserByID(id int) (model.User, error) {
	var user model.User
	result := initializer.PGDB.First(&user, id)
	return user, result.Error
}

func (r *PostgresUserRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	result := initializer.PGDB.Where("email = ?", email).First(&user)
	return user, result.Error
}

func (r *PostgresUserRepository) UpdateUser(id int, updatedUser model.User) (model.User, error) {
	var user model.User
	result := initializer.PGDB.First(&user, id)
	if result.Error != nil {
		return user, result.Error
	}

	result = initializer.PGDB.Model(&user).Updates(updatedUser)

	return user, result.Error
}

func (r *PostgresUserRepository) DeleteUser(id int) error {
	result := initializer.PGDB.Delete(&model.User{}, id)
	return result.Error
}
