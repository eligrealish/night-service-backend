package event

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	eventService EventService
)

type EventEndpoint struct {
	eventService EventService
}

func NewEventEndpoint() *EventEndpoint {
	eventService = *NewEventService()
	eventEndpoint := &EventEndpoint{eventService}
	return eventEndpoint
}

// the gin context is passed implictly when the router is registered
func (f EventEndpoint) GetEventWithParams(context *gin.Context) {
	list, err := eventService.GetEventParams(context)
	log.Println(err)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Internal Server Error",
		})
	} else {
		context.JSON(http.StatusOK, list)
	}
}

// the gin context is passed implictly when the router is registered
func (f EventEndpoint) GetEventByID(context *gin.Context) {
	event, err := eventService.GetEventByID(context)
	log.Println(err)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Internal Server Error",
		})
	} else {
		context.JSON(http.StatusOK, event)
	}
}
