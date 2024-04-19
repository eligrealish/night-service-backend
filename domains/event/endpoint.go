package event

import (
	"github.com/gin-gonic/gin"
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
	context.JSON(http.StatusOK, eventService.GetEventParams(context))
}

// the gin context is passed implictly when the router is registered
func (f EventEndpoint) GetEventByID(context *gin.Context) {
	context.JSON(http.StatusOK, eventService.GetEventByID(context))
}
