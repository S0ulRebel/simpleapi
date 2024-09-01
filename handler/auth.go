package handler

import (
	"net/http"
	"simple-api/errors"
	"simple-api/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func (h *AuthHandler) Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	err := c.BindJSON(&body)
	if err != nil {
		appErr := errors.NewErrorService().InternalServerError(err)
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	}

	tokenString, tokenStringErr := h.AuthService.Login(body.Email, body.Password)
	if tokenStringErr != nil {
		c.JSON(tokenStringErr.Code, gin.H{"error": tokenStringErr.Error()})
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}

func NewAuthHandler(
	authService service.AuthService,
) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}
