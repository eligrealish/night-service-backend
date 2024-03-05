package main

import (
	"log"
	"night-service-backend/routes"
)

func main() {
	//GET $DOMAIN/policy
	//GET $DOMAIN/event/$id
	//GET $DOMAIN/feed?location=$location
	//GET $DOMAIN/feed?location=$location&startDate=$startDate&endDate=$endDate

	log.Println("before router init")

	// Initialize Router
	routes.Initialize()
	log.Println("after router init")

}
