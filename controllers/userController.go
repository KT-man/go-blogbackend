package controllers

import (
	"context"
	"fmt"
	"go-backend/configs"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var usersCollection = configs.GetUsersCollection(configs.ConnectDB(), "users")
var validate = validator.New()

func PostCreateNewUser () gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		fmt.Println(&ctx)
		fmt.Println(usersCollection)
		fmt.Println(validate)
		
		defer cancel();
	}
}