package routes

import (
	"taptoeat-be/controllers"

	"github.com/gin-gonic/gin"
)

func DefineUser(r *gin.RouterGroup) {
	r.GET("/user", controllers.GetUser)
	r.POST("/user", controllers.CreateUser)
	r.PUT("/user/:id", controllers.PutUser)
	r.PATCH("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)
}
