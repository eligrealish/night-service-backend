package event

import (
	"context"
	"github.com/gin-gonic/gin"
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

func (e EventService) GetEventParams(context *gin.Context) Event {
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
		return Event{}
	}
	log.Println("dates not passed")
	GetEventByLocation(location)
	return Event{}
}

func GetEventByLocation(location string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("Night_Service").Collection("Event")
	log.Println(collection)
}

func GetEventByLocationAndDate(startDate string, endDate string, location string) {

}

func GetEventByID(context *gin.Context) {
	id := context.Param("id")
	log.Println(id)
}
