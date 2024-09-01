package service

import (
	"os"
	"simple-api/errors"
	"simple-api/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo repository.UserRepository
}

func (s AuthService) Login(email, password string) (string, *errors.AppError) {
	user, appErr := s.UserRepo.GetUserByEmail(email)
	if appErr != nil {
		return "", errors.NewErrorService().Unauthorized(errors.INVALID_EMAIL_OR_PASSWORD)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// Return a specific unauthorized error for incorrect password
		return "", errors.NewErrorService().Unauthorized(errors.INVALID_EMAIL_OR_PASSWORD)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return "", errors.NewErrorService().InternalServerError(err)
	}

	return tokenString, nil
}
