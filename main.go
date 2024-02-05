package main

import (
	"context"
	"fmt"
	"go-backend/configs"
	"go-backend/routes"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main () {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(configs.MongoEnvURI()).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
	  panic(err)
	}


	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key:"ping", Value:1}}).Err(); err != nil {
	  panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{"data": "New set"})
	})

	
	routes.UserRoutes(router)

	router.Run(":8080")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer func() {
    	if err = client.Disconnect(ctx); err != nil {
        panic(err)
    	}
	}()

	defer cancel()
}