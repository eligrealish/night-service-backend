package feed

import (
	"night-service-backend/domains/event" // assuming 'night-service-backend' is the module name
)

type Feed struct {
	Location string        `json:"location" bson:"location"`
	Events   []event.Event `json:"events" bson:"events"`
}
