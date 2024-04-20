package event

import (
	"context"
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

	// Determine if current time is before or after today's noon
	todayNoon := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location())
	var middayBefore, middayAfter time.Time

	if now.Before(todayNoon) {
		// If it's before noon, midday before is yesterday's noon
		middayBefore = todayNoon.AddDate(0, 0, -1)
		middayAfter = todayNoon
	} else {
		// If it's after noon, midday before is today's noon
		middayBefore = todayNoon
		middayAfter = todayNoon.AddDate(0, 0, 1)
	}

	// todo refactor to include last entry and to select 5 mins before that
	// Creating a filter to fetch events based on location and current date range
	filter := bson.M{
		"location": location,
		"endDate": bson.M{
			"$gte": middayBefore, // Event end date is greater than or equal to the last midday
			"$lt":  middayAfter,  // Event end date is less than the next midday
		},
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	var events []Event
	if err = cursor.All(context.TODO(), &events); err != nil {
		log.Fatal(err)
	}

	eventsDomain := List{
		Events: events,
	}

	// Here you can use eventsDomain as needed, e.g., print details, handle them in your application, etc.
	log.Println(eventsDomain)
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
