package routes

import (
	"github.com/gin-gonic/gin"
	"night-service-backend/domains/event"
	"night-service-backend/domains/feed"
	"night-service-backend/domains/policy"
)

// initializeRoutes initialize the routes of the API.
func initializeRoutes(router *gin.Engine) {
	policyEndpoint := policy.NewPolicyEndpoint()
	feedEndpoint := feed.NewFeedEndpoint()
	eventEndpoint := event.NewEventEndpoint()
	v1 := router.Group("/")
	{
		v1.GET("policy", policyEndpoint.GetPolicy)
		v1.GET("event", eventEndpoint.GetEvent)
		v1.GET("feed", feedEndpoint.GetFeed)
	}
}
