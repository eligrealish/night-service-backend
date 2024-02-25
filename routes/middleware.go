package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"night-service-backend/domains/policy"
)

var (
	policyService policy.PolicyService
)

type Middleware struct {
}

func NewMiddleware() *Middleware {
	middleware := &Middleware{}
	return middleware
}

func (m Middleware) CheckPolicyAccepted() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the value of the "X-App-Header" header
		headerValue := c.GetHeader("X-App-Header")
		log.Printf("header value : %s", headerValue)

		// Check if the header is set to "true"
		if headerValue != "true" {
			policyService = *policy.NewPolicyService()
			// If not, call policyEndpoint.GetPolicy and respond with the result
			c.JSON(http.StatusOK, policyService.GetPolicy())
			return
		}

		// If the header is set correctly, proceed to the next middleware/handler
		c.Next()
	}
}
