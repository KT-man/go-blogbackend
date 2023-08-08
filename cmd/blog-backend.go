package main

import (
	"go-backend/configs"

	"go-backend/routes"

	"github.com/gin-gonic/gin"
)

func main () {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{"data": "Testing idk what im doing"})
	})

	configs.ConnectDB()
	
	routes.UserRoutes(router)

	router.Run(":8080")
}