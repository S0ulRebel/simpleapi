package handler

import (
	"fmt"
	"net/http"
	"simple-api/errors"
	"simple-api/model"
	"simple-api/service"
	"simple-api/validation"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService service.UserService
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var body struct {
		FirstName string
		LastName  string
		Email     string
		Password  string
	}

	err := c.BindJSON(&body)
	if err != nil {
		appErr := errors.NewErrorService().BadRequest(errors.INVALID_REQUEST_BODY)
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	}

	user := model.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  body.Password,
	}
	if err := validation.Validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.INVALID_REQUEST_BODY})
		return
	}

	createdUser, createdUserErr := h.UserService.CreateUser(user)
	if createdUserErr != nil {
		c.JSON(createdUserErr.Code, gin.H{"error": createdUserErr.Error()})
		return
	}
	c.JSON(201, createdUser)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.UserService.GetUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.INVALID_ID})
		return
	}
	user, userErr := h.UserService.GetUserByID(id)
	if userErr != nil {
		c.JSON(500, gin.H{"error": userErr.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.INVALID_ID})
		return
	}
	var payload model.User

	payloadErr := c.BindJSON(&payload)
	if payloadErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.INVALID_REQUEST_BODY})
		return
	}
	fmt.Println(payload)
	updatedUser, updatedUserErr := h.UserService.UpdateUser(id, payload)
	if updatedUserErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": updatedUserErr.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.INVALID_ID})
		return
	}

	result := h.UserService.DeleteUser(id)
	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}
