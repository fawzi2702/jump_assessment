package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/this_is_iz/jump_server/pkg/handlers/user"
)

func SetupUserRouter(r *gin.Engine) {
	router := r.Group("/users")
	{
		router.GET("/", user.GetUsers)
		router.GET("/:user_id", user.GetUser)
		router.POST("/", user.CreateUser)
	}
}
