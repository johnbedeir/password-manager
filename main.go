package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Get user input for password name
	fmt.Print("Enter name for the password: ")
	var pwName string
	fmt.Scanln(&pwName)

	fmt.Print("Enter user name: ")
	var uName string
	fmt.Scanln(&uName)

	// Generate password
	rand.Seed(time.Now().UnixNano())
	pwLength := 16
	pwChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"
	var pw strings.Builder
	for i := 0; i < pwLength; i++ {
		pw.WriteByte(pwChars[rand.Intn(len(pwChars))])
	}

	// Connect to MongoDB and store password
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	collectionName := "passwords"
	collection := client.Database("passworddb").Collection(collectionName)

	type Password struct {
		Name     string `bson:"name"`
		Password string `bson:"password"`
	}

	_, err = collection.InsertOne(context.Background(), bson.M{"name": pwName, "username": uName, "password": pw.String()})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generated password:", pw.String())
	fmt.Println("Password stored in MongoDB database with name", pwName, "in collection", collectionName)
}
