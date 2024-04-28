package routes

import (
	"github.com/gin-gonic/gin"
	"night-service-backend/domains/event"
	"night-service-backend/domains/policy"
)

// initializeRoutes initialize the routes of the API.
func initializeRoutes(router *gin.Engine) {
	policyEndpoint := policy.NewPolicyEndpoint()
	eventEndpoint := event.NewEventEndpoint()
	endpoint := router.Group("/")
	{
		endpoint.GET("policy", policyEndpoint.GetPolicy)
		// todo review having two structs causes cocurrency issues
		endpoint.GET("event", eventEndpoint.GetEventWithParams)
		endpoint.GET("event/:id", eventEndpoint.GetEventByID)
	}
}
