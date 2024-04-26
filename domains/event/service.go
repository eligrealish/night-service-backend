package event

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"night-service-backend/utils"
	"time"
)

var (
	requestUtils utils.RequestUtils
)

type EventService struct {
}

func NewEventService() *EventService {
	eventService := &EventService{}
	requestUtils = *utils.RequestUtilsNewUtils()
	return eventService
}

func (e EventService) GetEventParams(context *gin.Context) List {
	log.Println("Event Handler")
	location, err := context.GetQuery("location")
	if err {
		log.Println(err)
		// should write 500 here
	}
	// Parse query string parameters
	if requestUtils.CheckRequestKeyPresent(context, "startDate") && requestUtils.CheckRequestKeyPresent(context, "endDate") {
		startDate, _ := context.GetQuery("startDate")
		endDate, _ := context.GetQuery("endDate")
		GetEventByLocationAndDate(startDate, endDate, location)
		log.Println("dates passed")
		return List{}
	}
	log.Println("dates not passed")
	GetEventByLocation(location)
	return List{}
}

// numbered in terms of dificuitly
// review impl here https://chat.openai.com/c/b06a5b18-13f7-4eb8-b017-b9266db04ac3
// todo retrives from mongo when the date is today 1
func GetEventByLocation(location string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("Night_Service").Collection("Event")

	// Get the current time
	now := time.Now()

	// Define the 9 AM time for the current day
	nineAm := time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, now.Location())

	var resultDate time.Time
	if now.Before(nineAm) {
		// If the current time is before 9 AM, set the result to the previous day
		resultDate = now.AddDate(0, 0, -1)
	} else {
		// If it is after 9 AM, set the result to the current day
		resultDate = now
	}

	filter := bson.D{{"dateOf", bson.D{{"eq", resultDate}}}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	// Iterate through the cursor
	var events []Event
	for cursor.Next(context.TODO()) {
		var event Event
		if err := cursor.Decode(event); err != nil {
			log.Fatal(err)
		}
		events = append(events, event)
	}
	log.Println(events)
	// Print the result
	fmt.Println("Resulting Date:", resultDate.Format("2006-01-02"))
}

// todo reterives from mongo via specific dates and location 2
func GetEventByLocationAndDate(startDate string, endDate string, location string) {

}

// todo reterives from mongo via ID 3
func (e EventService) GetEventByID(context *gin.Context) Event {
	id := context.Param("id")
	log.Println(id)
	return Event{}

}
