package handler

import (
	"net/http"
	"simple-api/errors"
	"simple-api/model"
	"simple-api/service"
	"simple-api/validation"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	PostService service.PostService
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	userId, idErr := c.Get("UserID")
	if !idErr {
		appErr := errors.NewErrorService().Unauthorized(errors.INVALID_ID)
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	}

	var body struct {
		UserID string
		Title  string
		Body   string
	}

	err := c.BindJSON(&body)
	if err != nil {
		appErr := errors.NewErrorService().BadRequest(errors.INVALID_REQUEST_BODY)
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	}

	post := model.Post{
		UserID: userId.(string),
		Title:  body.Title,
		Body:   body.Body,
	}

	if err := validation.Validate.Struct(post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.INVALID_REQUEST_BODY})
		return
	}

	createdPost, createdPostErr := h.PostService.CreatePost(post)
	if createdPostErr != nil {
		c.JSON(createdPostErr.Code, gin.H{"error": createdPostErr.Error()})
		return
	}
	c.JSON(201, createdPost)
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	posts, err := h.PostService.GetPosts()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, posts)
}

func (h *PostHandler) GetPostByID(c *gin.Context) {
	id := c.Param("id")
	post, err := h.PostService.GetPostByID(id)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, post)
}

func (h *PostHandler) GetPostsByUserID(c *gin.Context) {
	userID := c.Param("userID")
	posts, err := h.PostService.GetPostsByUserID(userID)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, posts)
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	userId, idErr := c.Get("UserID")
	if !idErr {
		appErr := errors.NewErrorService().Unauthorized(errors.INVALID_ID)
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	}
	id := c.Param("id")
	var body struct {
		Title string
		Body  string
	}

	err := c.BindJSON(&body)
	if err != nil {
		appErr := errors.NewErrorService().BadRequest(errors.INVALID_REQUEST_BODY)
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	}

	post := model.Post{
		ID:     id,
		UserID: userId.(string),
		Title:  body.Title,
		Body:   body.Body,
	}
	if err := validation.Validate.Struct(post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.INVALID_REQUEST_BODY})
		return
	}

	updatedPost, updatedPostErr := h.PostService.UpdatePost(id, post)
	if updatedPostErr != nil {
		c.JSON(updatedPostErr.Code, gin.H{"error": updatedPostErr.Error()})
		return
	}
	c.JSON(200, updatedPost)
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	id := c.Param("id")
	err := h.PostService.DeletePost(id)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
}

func NewPostHandler(postService service.PostService) *PostHandler {
	return &PostHandler{
		PostService: postService,
	}
}
