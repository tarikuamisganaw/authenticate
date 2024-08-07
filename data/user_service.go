package data

import (
	"context"
	"errors"
	"log"
	"tasker/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

func ConnectUserDB() error {
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

	userCollection = client.Database("taskdb").Collection("users")
	return nil
}

func CreateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	_, err = userCollection.InsertOne(ctx, user)
	return err
}

func AuthenticateUser(username, password string) (models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Searching for user with username:", username)
	log.Println("Searching for user with username:", password)
	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		log.Println("User not found or error in fetching user:", err)
		return user, errors.New("invalid username or password")
	}

	log.Println("User found, username:", user.Username)
	log.Println("Stored hashed password:", user.Password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println("Password mismatch:", err)
		return user, errors.New("invalid username or password")
	}

	log.Println("Password matched for user:", user.Username)
	return user, nil
}

// data/user_service.go
func GetAllUsers() ([]models.User, error) {
	var users []models.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return users, err
	}

	if err = cursor.All(ctx, &users); err != nil {
		return users, err
	}

	return users, nil
}
func GetUserByUsername(username string) (models.User, error) {
	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, errors.New("user not found")
		}
		return user, err
	}

	return user, nil
}

// Additional functions for user management go here
