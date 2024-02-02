package routes

import (
	"go-backend/controllers"

	"github.com/gin-gonic/gin"
)



func UserRoutes(router *gin.Engine) {
	// router.POST("/users/getAllUsers", controllers.PostGetAllUsers())
	router.POST("/users/createNewUser", controllers.PostCreateNewUser())
}