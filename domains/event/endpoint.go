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

// todo add error chcek, if given behaviour happens then error is returned and 500 is logged
// todo the err for 500 this will need to be considered with in all the services
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
	context.JSON(http.StatusOK, eventService.GetEventByID(context))
}
