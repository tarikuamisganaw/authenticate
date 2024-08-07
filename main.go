package main

import (
	"log"
	"tasker/data"
	"tasker/router"
)

func main() {
	// Connect to the task and user database
	err := data.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the task database:", err)
	}

	err = data.ConnectUserDB()
	if err != nil {
		log.Fatal("Failed to connect to the user database:", err)
	}

	// Set up the router
	r := router.SetupRouter()
	r.Run(":8080")
}
