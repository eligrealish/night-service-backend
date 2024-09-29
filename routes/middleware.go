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

			// required to stop request going to other route
			c.Abort()

			return
		}

		// If the header is set correctly, proceed to the next middleware/handler
		c.Next()
	}
}

func (m Middleware) CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
