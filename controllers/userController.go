package controllers

import (
	"fmt"
	responses "go-backend/api"
	"go-backend/configs"
	"go-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var usersCollection = configs.GetUsersCollection(configs.ConnectDB(), "users")
var validate = validator.New()

func PostCreateNewUser () gin.HandlerFunc {
	return func(c *gin.Context) {
		var NewUser models.NewUser
		
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
		newUserToAdd := models.User{
			UserId: primitive.NewObjectID(),
			Username: NewUser.Username,
			Password: NewUser.Password,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}
		fmt.Println(newUserToAdd)

		result, insertErr := usersCollection.InsertOne(c, newUserToAdd)
		
		if insertErr != nil {
			c.Error(insertErr)
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": insertErr.Error()})
			return 
		}
	
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
		c.JSON(http.StatusOK, gin.H{
			"Status": 200,
			"Message": "New User created",
		})
	}
}

func PostGetAllusers () gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}