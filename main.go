package main

import (
	"log"
	"night-service-backend/routes"
)

func main() {
	//GET $DOMAIN/policy
	//GET $DOMAIN/event/$id
	//GET $DOMAIN/event?location=$location
	//GET $DOMAIN/event?location=$location&startDate=$startDate&endDate=$endDate

	log.Println("before router init")

	// Initialize Router
	routes.Initialize()
	log.Println("after router init")

}
