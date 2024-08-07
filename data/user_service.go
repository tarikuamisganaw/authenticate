package data

import (
	"context"
	"errors"
	"log"
	"tasker/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

func InitializeUserService(client *mongo.Client, dbName string) {
	userCollection = client.Database(dbName).Collection("users")
}

func RegisterUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func AuthenticateUser(username, password string) (models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		log.Println("User not found or error in fetching user:", err)
		return user, errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println("Password mismatch:", err)
		return user, errors.New("invalid username or password")
	}

	return user, nil
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return users, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		if err = cursor.Decode(&user); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return users, err
	}

	return users, nil
}
