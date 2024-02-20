package routes

import (
	controllers "taptoeat-be/controllers"

	"github.com/gin-gonic/gin"
)

func DefineHelloWorld(r *gin.RouterGroup) {
	r.GET("/hello", controllers.Hello)
}
