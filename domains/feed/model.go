package feed

import (
	"night-service-backend/domains/event"
)

type Feed struct {
	Location string        `json:"location" bson:"location"`
	Events   []event.Event `json:"events" bson:"events"`
}
