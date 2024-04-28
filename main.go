package main

import (
	"night-service-backend/routes"
)

func main() {
	// TODO general clean up
	// TODO go through entire app and define way to remove repeated code e.g move to utils and other functions
	// TODO add more comments in increase clarity
	// TODO convert repo to run on Mac rather then WSL
	// TODO create new DB data and convert DB utils to work on Mac
	// TODO leave things like swagger till later
	// TODO consider adding tests like unit,regression or integration test before front end or after front end development
	// TODO consider if one or many test types should or should not be implemented
	// TODO remove ugly log statements
	// TODO review other TODOs
	// TODO review if ./routes/router.go should be moved to ./main.go
	// TODO review if error should be moved to model or utils
	// TODO review if if else for 500 or 200 should be moved to utils
	// TODO review logging can be improved

	// Initialize Router
	routes.Initialize()

}
