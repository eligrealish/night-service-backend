package event

import (
	"time"
)

type Event struct {
	EventName   string    `json:"eventName" bson:"eventName"`
	Venue       string    `json:"venue" bson:"venue"`
	Location    string    `json:"location" bson:"location"`
	StartDate   time.Time `json:"startDate" bson:"startDate"`
	EndDate     time.Time `json:"endDate" bson:"endDate"`
	Images      []string  `json:"images" bson:"images"`
	Description string    `json:"description" bson:"description"`
}