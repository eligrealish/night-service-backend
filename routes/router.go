package routes

import (
	"github.com/gin-gonic/gin"
)

// Initializes and run the router.
func Initialize() {
	router := gin.Default()

	middleware := Middleware{}
	router.Use(middleware.CheckPolicyAccepted())
	initializeRoutes(router)

	err := router.Run()
	if err != nil {
		return
	}
}
