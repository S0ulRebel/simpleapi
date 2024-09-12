package postgres

import (
	"simple-api/errors"
	"simple-api/initializer"
	"simple-api/model"
)

type PostgresPostRepository struct {
}

func (r *PostgresPostRepository) CreatePost(Post model.Post) (model.Post, *errors.AppError) {
	err := initializer.PGDB.Exec("INSERT INTO posts (id, user_id, title, body) VALUES ($1, $2, $3, $4)", Post.ID, Post.UserID, Post.Title, Post.Body).Error
	if err != nil {
		return Post, errors.NewErrorService().InternalServerError(err)
	}
	return Post, nil
}

func (r *PostgresPostRepository) GetPosts() ([]model.Post, *errors.AppError) {
	var posts []model.Post
	err := initializer.PGDB.Find(&posts).Error
	if err != nil {
		return posts, errors.NewErrorService().InternalServerError(err)
	}
	return posts, nil
}

func (r *PostgresPostRepository) GetPostByID(id string) (model.Post, *errors.AppError) {
	var post model.Post
	err := initializer.PGDB.Where("id = ?", id).First(&post).Error
	if err != nil {
		return post, errors.NewErrorService().InternalServerError(err)
	}
	return post, nil
}

func (r *PostgresPostRepository) GetPostsByUserID(userID string) ([]model.Post, *errors.AppError) {
	var posts []model.Post
	err := initializer.PGDB.Where("user_id = ?", userID).Find(&posts).Error
	if err != nil {
		return posts, errors.NewErrorService().InternalServerError(err)
	}
	return posts, nil
}

func (r *PostgresPostRepository) UpdatePost(id string, updatedPost model.Post) (model.Post, *errors.AppError) {
	err := initializer.PGDB.Model(&model.Post{}).Where("id = ?", id).Updates(updatedPost).Error
	if err != nil {
		return updatedPost, errors.NewErrorService().InternalServerError(err)
	}
	return updatedPost, nil
}

func (r *PostgresPostRepository) DeletePost(id string) *errors.AppError {
	err := initializer.PGDB.Where("id = ?", id).Delete(&model.Post{}).Error
	if err != nil {
		return errors.NewErrorService().InternalServerError(err)
	}
	return nil
}
