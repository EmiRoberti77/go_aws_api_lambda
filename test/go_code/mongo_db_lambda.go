package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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

// Handler is the Lambda handler function.
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    // Your MongoDB connection string, database, and collection names.
    connectionString := "mongodb+srv://vetapp:vetapp123@cluster0.uwhi5uh.mongodb.net/?retryWrites=true&w=majority"
    dbName := "vetapp"
    collectionName := "equines"

    // Connect to the MongoDB server.
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
    if err != nil {
        fmt.Printf("Error connecting to MongoDB: %v\n", err)
        return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Internal Server Error"}, err
    }
    defer client.Disconnect(ctx)

    // Select the database and collection.
    database := client.Database(dbName)
    collection := database.Collection(collectionName)

    // Create an Equine object.
    equine := Equine{
        UELN:           2999909000,
        Type:           "horse type A2",
        Color:          "pink",
        Vaccination:    "completed",
        VaccinationDate: "1/1/2024",
    }

    // Insert the Equine object into the collection.
    _, err = collection.InsertOne(ctx, equine)
    if err != nil {
        fmt.Printf("Error inserting document: %v\n", err)
        return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Internal Server Error"}, err
    }

    return events.APIGatewayProxyResponse{StatusCode: 200, Body: "Equine document inserted successfully"}, nil
}

func main() {
    lambda.Start(Handler)
}