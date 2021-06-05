package router

import (
	"ahead/app/controller/user"

	"github.com/gin-gonic/gin"
)

func RegisterApi() *gin.Engine {

	route := gin.Default()

	rUser := route.Group("/api/v1/user")
	{
		userCtrl := new(user.UserController)
		rUser.GET("/", userCtrl.Index)
		rUser.GET("/:id", userCtrl.Show)
		rUser.POST("/", userCtrl.Create)
		rUser.PUT("/", userCtrl.Update)
		rUser.DELETE("/", userCtrl.Delete)
	}

	return route
}
