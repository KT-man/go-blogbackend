package routes

import (
	"go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/user/createNewUser", controllers.PostCreateNewUser())
}