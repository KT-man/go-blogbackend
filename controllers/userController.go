package controllers

import (
	"context"
	"fmt"
	responses "go-backend/api"
	"go-backend/configs"
	"go-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var usersCollection = configs.GetUsersCollection(configs.ConnectDB(), "users")
var validate = validator.New()

func PostCreateNewUser () gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		fmt.Println(ctx)
		var NewUser models.NewUser
		defer cancel()

		// Validate that request body fits the &user struct
		
		err := c.BindJSON(&NewUser); if err != nil {
			fmt.Println("Failed")
			c.JSON(http.StatusBadRequest, responses.ApiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
		}

		 //Use the validator library to validate required fields
        if validationErr := validate.Struct(&NewUser); validationErr != nil {
			fmt.Println("This is error",validationErr)		
            c.JSON(http.StatusBadRequest, responses.ApiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }

		// Handle valid request and add single document to collection
		fmt.Println(&NewUser)
		fmt.Println(NewUser.Username)
		fmt.Println("It's workign!")
	}
}

func PostGetAllusers () gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}