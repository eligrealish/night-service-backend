package main

import (
	"backend/model"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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
	var result bson.M

	// Find the document
	err = collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	bsonBytes, _ := bson.Marshal(result)

	// Print the result
	fmt.Printf("Found document: %T (%s)", result, result)

	var messageModel *model.Message

	// Convert the BSON document to a Message struct
	err = bson.Unmarshal(bsonBytes, &messageModel)
	if err != nil {
		fmt.Println("Error unmarshaling BSON:", err)
		return
	}

	if err = client.Disconnect(ctx); err != nil {
		panic(err)
	}

	fmt.Printf("mapped Struct: %T (%s)", messageModel, messageModel)

	// TODO https://chat.openai.com/c/296f6a11-de68-462f-97b5-190d8661831b wtf is this
	// cross reference with these
	// TODO https://www.mongodb.com/docs/drivers/go/current/fundamentals/bson/
	// TODO https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/query-document/
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}
