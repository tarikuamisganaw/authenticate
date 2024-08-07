package data

import (
	"context"
	"errors"
	"tasker/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var taskCollection *mongo.Collection

func ConnectDB() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	taskCollection = client.Database("taskdb").Collection("tasks")
	return nil
}

func GetTaskByID(idStr string) (models.Task, error) {
	var task models.Task

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := taskCollection.FindOne(ctx, bson.M{"_id": idStr}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return task, errors.New("task not found")
		}
		return task, err
	}

	return task, nil
}

// Additional CRUD functions (CreateTask, UpdateTask, DeleteTask, GetTasks) go here
func CreateTask(task models.Task) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return taskCollection.InsertOne(ctx, task)
}

// GetTasks retrieves all tasks from the MongoDB collection
func GetTasks() ([]models.Task, error) {
	var tasks []models.Task

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := taskCollection.Find(ctx, bson.M{})
	if err != nil {
		return tasks, err
	}

	for cursor.Next(ctx) {
		var task models.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}

	return tasks, cursor.Err()
}

// GetTaskByID retrieves a task by its ID from the MongoDB collection

// UpdateTask updates an existing task in the MongoDB collection
func UpdateTask(id string, task models.Task) (*mongo.UpdateResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return taskCollection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.D{
		{"$set", task},
	})
}

// DeleteTask deletes a task from the MongoDB collection by its ID
func DeleteTask(id string) (*mongo.DeleteResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return taskCollection.DeleteOne(ctx, bson.M{"_id": objectID})
}
