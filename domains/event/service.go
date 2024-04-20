package event

import (
	"github.com/gin-gonic/gin"
	"log"
	"night-service-backend/utils"
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
	//	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	//	collection := client.Database("Night_Service").Collection("Event")

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
