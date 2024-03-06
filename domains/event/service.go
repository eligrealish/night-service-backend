package event

import (
	"github.com/gin-gonic/gin"
	"log"
	"night-service-backend/utils"
)

var (
	utils Utils
)


type EventService struct {
}

func NewEventService() *EventService {
	eventService := &EventService{}
	utils := *NewUtils()
	return eventService
}

func (e EventService) GetEventHandler(context *gin.Context) Event {
	log.Println("Event Handler")
	// Parse query string parameters
	if utils.checkRequestKeyPresent(context,"startDate") && utils.checkRequestKeyPresent(context,"endDate")
	return Event{}
}
