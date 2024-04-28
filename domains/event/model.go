package event

import "time"

type List struct {
	Country string  `json:"country" bson:"country"`
	City    string  `json:"city" bson:"city"`
	DateOf  string  `json:"dateOf" bson:"location"`
	Events  []Event `json:"events,omitempty" bson:"events"` // JSON field is optional
}
type Event struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	EventName   string    `json:"eventName" bson:"eventName"`
	Venue       string    `json:"venue" bson:"venue"`
	StartDate   time.Time `json:"startDate" bson:"startDate"`
	EndDate     time.Time `json:"endDate" bson:"endDate"`
	Images      []string  `json:"images" bson:"images"`
	Description string    `json:"description" bson:"description"`
	DateOf      string    `json:"dateOf" bson:"dateOf"`
	Country     string    `json:"country" bson:"country"`
	City        string    `json:"city" bson:"city"`
}
