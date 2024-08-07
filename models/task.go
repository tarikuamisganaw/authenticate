package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	DueDate     string             `json:"due_date"`
	Status      string             `json:"status"`
}
