package routes

import (
	"github.com/gin-gonic/gin"
	"log"
)

// Initializes and run the router.
func Initialize() {
	log.Println("init func")
	//TODO add to gin New later
	router := gin.Default()
	log.Println("router struct created")

	log.Println("before routes registered")
	initializeRoutes(router)
	log.Println("after routes registered")

	err := router.Run()
	log.Println("running router!")
	if err != nil {
		return
	}
}
