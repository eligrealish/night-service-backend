package event

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	eventService EventService
)

type EventEndpoint struct {
	EventService EventService
}

func NewEventEndpoint() *EventEndpoint {
	eventService = *NewEventService()
	eventEndpoint := &EventEndpoint{eventService}
	return eventEndpoint
}

// the gin context is passed implictly when the router is registered
func (e EventEndpoint) GetEvent(context *gin.Context) {
	context.JSON(http.StatusOK, eventService.GetEventHandler(context))
}
