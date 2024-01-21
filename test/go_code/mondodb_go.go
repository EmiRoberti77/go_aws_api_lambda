package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Equine represents the data structure for the object to be inserted.
type Equine struct {
    UELN           int    `bson:"ueln"`
    Type           string `bson:"type"`
    Color          string `bson:"color"`
    Vaccination    string `bson:"vaccination"`
    VaccinationDate string `bson:"vaccinationDate"`
}

func main() {
    // MongoDB connection string.
    connectionString := "mongodb://your_username:your_password@your_mongodb_host:your_mongodb_port/?authSource=your_auth_database"

    // MongoDB database and collection names.
    dbName := "your_database_name"
    collectionName := "equine"

    // Create a context with a timeout for database operations.
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Connect to the MongoDB server.
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
    if err != nil {
        fmt.Printf("Error connecting to MongoDB: %v\n", err)
        return
    }

    // Select the database and collection.
    database := client.Database(dbName)
    collection := database.Collection(collectionName)

    // Create an Equine object.
    equine := Equine{
        UELN:           23456677,
        Type:           "horse type A2",
        Color:          "black",
        Vaccination:    "completed",
        VaccinationDate: "22/1/2024",
    }

    // Insert the Equine object into the collection.
    _, err = collection.InsertOne(ctx, equine)
    if err != nil {
        fmt.Printf("Error inserting document: %v\n", err)
        return
    }

    fmt.Println("Equine document inserted successfully!")
}