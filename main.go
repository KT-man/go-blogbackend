package main

import (
	"go-backend/configs"

	"go-backend/routes"

	"github.com/gin-gonic/gin"
)

func main () {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{"data": "New set"})
	})

	configs.ConnectDB()
	
	routes.UserRoutes(router)

	router.Run(":8080")
}