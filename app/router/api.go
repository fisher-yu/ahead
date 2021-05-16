package router

import (
	"ahead/app/controller/user"

	"github.com/gin-gonic/gin"
)

func RegisterApi() *gin.Engine {

	route := gin.Default()

	rUser := route.Group("/api/v1/user")
	{
		rUser.GET("/", user.UserController{}.Index)
		rUser.GET("/:id", user.UserController{}.Show)
		rUser.POST("/", user.UserController{}.Create)
		rUser.PUT("/", user.UserController{}.Update)
		rUser.DELETE("/", user.UserController{}.Delete)
	}

	return route
}
