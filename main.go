package main

import (
	"go-gin-gorm/config"
	"go-gin-gorm/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	// r := gin.Default()
	// r.Run()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8082")
}
