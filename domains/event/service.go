package event

import (
	"context"
	"errors"
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

func (e EventService) GetEventParams(context *gin.Context) (List, error) {
	log.Println("Event Handler")
	countryCode, err := context.GetQuery("countryCode")
	countryCode = strings.ToUpper(countryCode)
	city, err := context.GetQuery("city")
	city = charUtils.UppercaseFirst(city)
	if !err {
		return List{}, errors.New("missing parameters")
	}
	// Parse query string parameters
	if requestUtils.CheckRequestKeyPresent(context, "startDate") && requestUtils.CheckRequestKeyPresent(context, "endDate") {
		log.Println("dates passed")

		startDate, _ := context.GetQuery("startDate")
		endDate, _ := context.GetQuery("endDate")
		return GetEventByLocationAndDate(startDate, endDate, countryCode, city)
	}
	log.Println("dates not passed")
	return GetEventByLocation(countryCode, city)
}

func GetEventByLocation(countryCode string, city string) (List, error) {
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

	list, err := GetEventByDBQuery(filter, countryCode, city)
	if err != nil {
		return List{}, err
	}
	list.DateOf = resultDate.Format("2006-01-02")

	return list, nil
}

func GetEventByLocationAndDate(startDate string, endDate string, countryCode string, city string) (List, error) {
	filter := bson.D{
		{"dateOf", bson.D{
			{"$gte", startDate},
			{"$lte", endDate},
		}},
		{"countryCode", countryCode},
		{"city", city},
	}
	return GetEventByDBQuery(filter, countryCode, city)
}

func GetEventByDBQuery(filter bson.D, countryCode string, city string) (List, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("Night_Service").Collection("Event")

	log.Printf("filter : %s", filter)
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Println("no database response")
		return List{}, errors.New("no database response")
	}

	var eventsPtr []*Event
	for cursor.Next(context.Background()) {
		var event Event
		err := cursor.Decode(&event)
		if err != nil {
			return List{}, err

		}
		eventsPtr = append(eventsPtr, &event)
		log.Println("event : %s", event)
	}

	// Check if no events were found
	// is suitable to do here as the cursor logic will skip so expense increase is relatively minimal
	if len(eventsPtr) == 0 {
		log.Println("No events found matching the filter")
		return List{}, errors.New("no events found matching the filter")
	}

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

	log.Println(list)
	return list, nil
}

// todo reterives from mongo via ID 3
func (e EventService) GetEventByID(context *gin.Context) Event {
	id := context.Param("id")
	log.Println(id)
	return Event{}

}
