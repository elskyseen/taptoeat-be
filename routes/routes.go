package routes

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		DefineHelloWorld(v1)
		DefineUser(v1)
	}
}
