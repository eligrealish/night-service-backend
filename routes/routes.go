package routes

import (
	"github.com/gin-gonic/gin"
	"night-service-backend/domains/policy"
)

//reference for this design
//https://github.com/gustavoopedrosa/go-gin-sqlite/blob/main/router/
//https://github.com/eddycjy/go-gin-example	/tree/master/routers/api

// initializeRoutes initialize the routes of the API.
func initializeRoutes(router *gin.Engine) {
	policyEndpoint := policy.NewPolicyEndpoint()
	v1 := router.Group("/api/v1")
	{
		v1.GET("policy", policyEndpoint.GetPolicy)
		v1.GET("test", policyEndpoint.GetTest)
		//v1.POST(openingRoute, handler.CreateOpeningHandler)
		//v1.DELETE(openingRoute, handler.DeleteOpeningHandler)
		//v1.PUT(openingRoute, handler.UpdateOpeningHandler)
		//v1.GET("/openings", handler.ListOpeningsHandler)
	}
}

//TODO add do header that must be sent as true to not show privacy policy
//TODO add an endpoint that show policy not from header
//TODO give header name like x-app-header, x meaning no specific and app-header because its true if you are requesting the app as the policy is "outside" in the app
