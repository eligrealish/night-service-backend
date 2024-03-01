package policy

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

//TODO hardcode and don't use model for testing

type PolicyService struct {
}

func NewPolicyService() *PolicyService {
	instance := &PolicyService{}
	return instance
}

func (s PolicyService) GetPolicy() Message {
	log.Println("logging Policy Instance")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("Night_Service").Collection("Privacy_Notice")
	filter := bson.D{}

	// Create a variable in which to store the result
	//var result bson.M

	var message Message
	// Find the document
	err = collection.FindOne(context.Background(), filter).Decode(&message)
	fmt.Printf("thing is %s", message)

	if err = client.Disconnect(ctx); err != nil {
		panic(err)
	}
	return message
}
