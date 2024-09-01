package middleware

import (
	"fmt"
	"os"
	"simple-api/errors"
	"simple-api/model"
	"simple-api/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, tokenStringErr := c.Cookie("Authorization")
		if tokenStringErr != nil {
			handleAuthError(c, "No token provided")
			return
		}

		token, tokenErr := parseToken(tokenString)
		if tokenErr != nil {
			handleAuthError(c, "Unexpected token signing method")
			return
		}

		claims, valid := token.Claims.(jwt.MapClaims)
		if !valid || !token.Valid {
			handleAuthError(c, "Invalid token or claims")
			return
		}

		if claimsErr := validateClaims(claims); claimsErr != nil {
			handleAuthError(c, claimsErr.Error())
			return
		}

		user, userErr := getUserByClaims(claims, userService)
		if userErr != nil {
			handleAuthError(c, "Error extracting ID claim")
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func handleAuthError(c *gin.Context, message string) {
	appErr := errors.NewErrorService().Unauthorized(message)
	c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
	c.Abort()
}

func parseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
}

func validateClaims(claims jwt.MapClaims) error {
	exp, ok := claims["exp"].(float64)
	if !ok {
		return fmt.Errorf("Error extracting expiration claim")
	}
	if exp < float64(time.Now().Unix()) {
		return fmt.Errorf("Token expired")
	}

	_, ok = claims["ID"].(float64)
	if !ok {
		return fmt.Errorf("Error extracting ID claim")
	}

	return nil
}

func getUserByClaims(claims jwt.MapClaims, userService service.UserService) (model.User, *errors.AppError) {
	id, ok := claims["ID"].(float64)
	if !ok {
		return model.User{}, errors.NewErrorService().Unauthorized("")
	}
	return userService.GetUserByID(int(id))
}
