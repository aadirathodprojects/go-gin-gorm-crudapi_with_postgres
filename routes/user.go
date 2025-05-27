package routes

import (
	"go-gin-gorm/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controller.CreateUser)
	router.GET("/users", controller.GetUsers)
	router.PUT("/updateuser/:id", controller.UpdateUser)
	router.DELETE("/deleteuser/:id", controller.DeleteUser)
}
