package service

import (
	"fmt"
	"simple-api/errors"
	"simple-api/model"
	"simple-api/repository"

	"github.com/google/uuid"
)

type PostService struct {
	PostRepo repository.PostRepository
}

func (s *PostService) CreatePost(post model.Post) (model.Post, *errors.AppError) {
	post.ID = uuid.New().String()
	fmt.Println(post)
	return s.PostRepo.CreatePost(post)
}

func (s *PostService) GetPosts() ([]model.Post, *errors.AppError) {
	return s.PostRepo.GetPosts()
}

func (s *PostService) GetPostByID(id string) (model.Post, *errors.AppError) {
	return s.PostRepo.GetPostByID(id)
}

func (s *PostService) GetPostsByUserID(userID string) ([]model.Post, *errors.AppError) {
	return s.PostRepo.GetPostsByUserID(userID)
}

func (s *PostService) UpdatePost(id string, updatedPost model.Post) (model.Post, *errors.AppError) {
	return s.PostRepo.UpdatePost(id, updatedPost)
}

func (s *PostService) DeletePost(id string) *errors.AppError {
	return s.PostRepo.DeletePost(id)
}
