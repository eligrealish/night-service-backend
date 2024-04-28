package event

import "time"

type List struct {
	CountryCode string  `json:"countryCode" bson:"countryCode"`
	City        string  `json:"city" bson:"city"`
	DateOf      string  `json:"dateOf,omitempty" bson:"location"` // JSON field is optional
	Events      []Event `json:"events" bson:"events"`
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
	CountryCode string    `json:"countryCode" bson:"countryCode"`
	City        string    `json:"city" bson:"city"`
}
