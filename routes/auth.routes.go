package routes

import (
	"taptoeat-be/controllers"

	"github.com/gin-gonic/gin"
)

func DefineAuth(c *gin.RouterGroup) {
	c.POST("/auth/signup",controllers.Signup)
	c.POST("/auth/signin",controllers.SignIn)
}
