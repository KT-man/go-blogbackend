package configs

import "go.mongodb.org/mongo-driver/mongo"

func GetUsersCollection (client *mongo.Client, collectionName string) *mongo.Collection {
	usersCollection := client.Database("blog-cluster").Collection(collectionName)
	return usersCollection
}