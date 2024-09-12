package repository

import (
	"simple-api/errors"
	"simple-api/model"
)

type PostRepository interface {
	CreatePost(Post model.Post) (model.Post, *errors.AppError)
	GetPosts() ([]model.Post, *errors.AppError)
	GetPostByID(id string) (model.Post, *errors.AppError)
	GetPostsByUserID(userID string) ([]model.Post, *errors.AppError)
	UpdatePost(id string, updatedPost model.Post) (model.Post, *errors.AppError)
	DeletePost(id string) *errors.AppError
}
