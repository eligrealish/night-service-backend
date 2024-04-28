package event

import "time"

type List struct {
	Location string  `json:"location" bson:"location"`
	DateOf   string  `json:"dateOf" bson:"location"`
	Events   []Event `json:"events,omitempty" bson:"events"` // JSON field is optional
}
type Event struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	EventName   string    `json:"eventName" bson:"eventName"`
	Venue       string    `json:"venue" bson:"venue"`
	Location    string    `json:"location" bson:"location"`
	StartDate   time.Time `json:"startDate" bson:"startDate"`
	EndDate     time.Time `json:"endDate" bson:"endDate"`
	Images      []string  `json:"images" bson:"images"`
	Description string    `json:"description" bson:"description"`
	DateOf      string    `json:"dateOf" bson:"dateOf"`
}
