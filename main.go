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

func getPassName() string {
	fmt.Print("Enter name for the password: ")
	var pwName string
	fmt.Scanln(&pwName)
	return pwName
}

func getUserName() string {
	fmt.Print("Enter user name: ")
	var uName string
	fmt.Scanln(&uName)
	return uName
}

func generatePass() string {
	rand.Seed(time.Now().UnixNano())
	pwLength := 16
	pwChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"
	var pw strings.Builder
	for i := 0; i < pwLength; i++ {
		pw.WriteByte(pwChars[rand.Intn(len(pwChars))])
	}
	return pw.String()
}

// Connect to Database and store the data
func databaseConnect(pwName, uName, password string) (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	collectionName := "passwords"
	collection := client.Database("passworddb").Collection(collectionName)

	type Password struct {
		Name     string `bson:"name"`
		Password string `bson:"password"`
	}

	encryptedPassword := encrypt(password)
	_, err = collection.InsertOne(context.Background(), bson.M{"name": pwName, "username": uName, "password": encryptedPassword})
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func printPass(pwName, password string, collectionName string) {
	fmt.Println("Generated password:", password)
	fmt.Println("Password stored in MongoDB database with name", pwName, "in collection", collectionName)
}

func main() {
	pwName := getPassName()
	uName := getUserName()
	password := generatePass()
	collection, err := databaseConnect(pwName, uName, password)
	if err != nil {
		log.Fatal(err)
	}
	printPass(pwName, password, collection.Name())
}
