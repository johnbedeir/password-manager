package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getPasswordName() string {
	fmt.Print("Enter the name of the password: ")
	var pwName string
	fmt.Scanln(&pwName)
	return pwName
}

func fetchPassword(pwName string) (string, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return "", err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("passworddb").Collection("passwords")
	var result bson.M
	err = collection.FindOne(context.Background(), bson.M{"name": pwName}).Decode(&result)
	if err != nil {
		return "", err
	}
	return result["password"].(string), nil
}

func main() {
	pwName := getPasswordName()
	encryptedPassword, err := fetchPassword(pwName)
	if err != nil {
		log.Fatal(err)
	}
	password := decrypt(encryptedPassword)
	fmt.Println("Decrypted password for", pwName, "is:", password)
}
