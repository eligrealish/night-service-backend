package main

import (
	"log"
	"night-service-backend/routes"
)

// TODO start a github notes, authenticate with an SSH key
// TODO make a TODO.go? to list things you haven't got a markdown repo yet
// TODO note down how to pull in dependencies "go get {$project_name} ."

func main() {
	// TODO middleware broken, likely due to pointers in router package

	// TODO next Mongo backing up -> https://chat.openai.com/c/70c1bbc2-62d3-435a-84b4-25d89e2aaed2

	//TODO cross reference  error with pointer question

	//TODO review other debug methods then just log breakpoints
	// info log, can this be added into Gin?
	log.Println("before router init")

	// Initialize Router
	routes.Initialize()
	log.Println("after router init")

}

//TODO below done you need to use structs, they encapslate data and passing data can be slow.
//TODO read this to understand how/why/when to write golang structs vs funcs. Possibly cross ref with C and C++ in chat gpt? https://chat.openai.com/c/0010bcc5-38c3-4186-ac35-6017ea1e0685
