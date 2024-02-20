package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"night-service-backend/domain/models"
	"time"
)

// TODO start a github notes, authenticate with an SSH key
// TODO make a TODO.go? to list things you haven't got a markdown repo yet
// TODO note down how to pull in dependencies "go get {$project_name} ."

func main() {
	router := gin.Default()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	collection := client.Database("privacy_notice").Collection("message")

	filter := bson.D{}

	// Create a variable in which to store the result
	//var result bson.M

	var message models.Message
	// Find the document
	err = collection.FindOne(context.Background(), filter).Decode(&message)

	fmt.Printf("thing is %s", message)

	if err = client.Disconnect(ctx); err != nil {
		panic(err)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, message)
	})

	router.Run(":8080")

}
