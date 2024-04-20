package main

import (
	"log"
	"night-service-backend/domains/event"
)

func main() {
	//GET $DOMAIN/policy
	//GET $DOMAIN/event/$id
	//GET $DOMAIN/event?location=$location
	//GET $DOMAIN/event?location=$location&startDate=$startDate&endDate=$endDate

	eventService := event.NewEventService()

	thing := eventService.GetEventByLocation("brighton")

	log.Println("before router init")

	// Initialize Router
	//routes.Initialize()
	log.Println("after router init")

}
