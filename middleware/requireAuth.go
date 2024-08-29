package middleware

import (
	"fmt"
	"net/http"
	"os"
	"simple-api/model"
	"simple-api/initializer"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp, ok := claims["exp"].(float64)
		if !ok {
			fmt.Println("Error extracting exp claim")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if exp < float64(time.Now().Unix()) {
			fmt.Println("Token has expired")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		id, ok := claims["sub"].(float64)
		if !ok {
			fmt.Println("Error extracting id claim")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user model.User
		if err := initializer.PGDB.First(&user, int(id)).Error; err != nil {
			fmt.Println("Error fetching user from database:", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)
		c.Next()
	} else {
		fmt.Println("Invalid token or claims")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
