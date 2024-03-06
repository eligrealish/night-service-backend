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

func (e EventEndpoint) GetEvent(context *gin.Context) {
	context.JSON(http.StatusOK, Event{})
}
