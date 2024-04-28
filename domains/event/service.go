package event

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"night-service-backend/utils"
	"strings"
	"time"
)

var (
	requestUtils utils.RequestUtils
	charUtils    utils.CharUtils
)

type EventService struct {
}

func NewEventService() *EventService {
	eventService := &EventService{}
	requestUtils = *utils.RequestUtilsNewUtils()
	charUtils = *utils.NewCharUtils()
	return eventService
}

func (e EventService) GetEventParams(context *gin.Context) List {
	log.Println("Event Handler")
	countryCode, err := context.GetQuery("countryCode")
	countryCode = strings.ToUpper(countryCode)
	city, err := context.GetQuery("city")
	city = charUtils.UppercaseFirst(city)
	if err {
		log.Println(err)
		// should write 500 here
	}
	// Parse query string parameters
	if requestUtils.CheckRequestKeyPresent(context, "startDate") && requestUtils.CheckRequestKeyPresent(context, "endDate") {
		startDate, _ := context.GetQuery("startDate")
		endDate, _ := context.GetQuery("endDate")
		GetEventByLocationAndDate(startDate, endDate, countryCode, city)
		log.Println("dates passed")
		return List{}
	}
	log.Println("dates not passed")
	return GetEventByLocation(countryCode, city)
}

// numbered in terms of dificuitly
// review impl here https://chat.openai.com/c/b06a5b18-13f7-4eb8-b017-b9266db04ac3
// todo retrives from mongo when the date is today 1
func GetEventByLocation(countryCode string, city string) List {

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

	filter := bson.D{{"dateOf", resultDate.Format("2006-01-02")},
		{"countryCode", countryCode},
		{"city", city},
	}
	log.Printf("filter : %s", filter)
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var eventsPtr []*Event
	for cursor.Next(context.Background()) {
		var event Event
		err := cursor.Decode(&event)
		if err != nil {
			log.Println(err)

		}
		eventsPtr = append(eventsPtr, &event)
		log.Println(event)
	}
	log.Println(eventsPtr)
	var list List
	events := make([]Event, 0, len(eventsPtr))
	// Convert []*Event to []Event
	for _, e := range eventsPtr {
		if e != nil {
			events = append(events, *e) // Dereference the pointer and append the struct
		}
	}
	list.Events = events
	list.CountryCode = countryCode
	list.City = city
	list.DateOf = resultDate.Format("2006-01-02")
	log.Println(list)
	return list
}

// todo reterives from mongo via specific dates and location 2
func GetEventByLocationAndDate(startDate string, endDate string, countryCode string, city string) {

}

// todo reterives from mongo via ID 3
func (e EventService) GetEventByID(context *gin.Context) Event {
	id := context.Param("id")
	log.Println(id)
	return Event{}

}
