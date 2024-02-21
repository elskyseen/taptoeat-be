package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Middleware(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message":     "Invalid Request",
			"redirectUrl": "/login",
			"code":        http.StatusUnauthorized,
		})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":     "Invalid Request",
				"redirectUrl": "/login",
				"code":        http.StatusUnauthorized,
			})
			return
		}
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Request",
			"code":    http.StatusUnauthorized,
		})
		return
	}
}
