package main

import (
	"context"
	"log"
	"tasker/data"
	"tasker/router"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	data.InitializeUserService(client, "task_manager")
	data.InitializeTaskService(client, "task_manager")

	r := router.SetupRouter()
	r.Run(":8080")
}
