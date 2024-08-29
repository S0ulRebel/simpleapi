package handler

import (
	"net/http"
	"simple-api/model"
	"simple-api/service"
	"strconv"

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

	c.BindJSON(&body)

	user := model.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  body.Password,
	}

	createdUser, err := h.UserService.CreateUser(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	user, err := h.UserService.GetUserByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var payload model.User

	c.BindJSON(&payload)

	updatedUser, err := h.UserService.UpdateUser(id, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result := h.UserService.DeleteUser(id)
	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// func Login(c *gin.Context) {
// 	var body struct {
// 		Email    string
// 		Password string
// 	}
// 	c.BindJSON(&body)

// 	var user model.User
// 	result := initializer.PGDB.Where("email = ?", body.Email).First(&user)
// 	if result.Error != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
// 		return
// 	}

// 	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
// 		return
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"sub": user.ID,
// 		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
// 	})

// 	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.SetSameSite(http.SameSiteNoneMode)
// 	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
// 	c.JSON(http.StatusOK, gin.H{"token": tokenString})
// }

// func Logout(c *gin.Context) {
// 	c.SetSameSite(http.SameSiteNoneMode)
// 	c.SetCookie("Authorization", "", -1, "", "", false, true)
// 	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
// }

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}
