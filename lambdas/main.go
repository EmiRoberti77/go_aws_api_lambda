package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type Equine struct {
    UELN           int    `bson:"ueln"`
    Type           string `bson:"type"`
    Color          string `bson:"color"`
    Vaccination    string `bson:"vaccination"`
    VaccinationDate string `bson:"vaccinationDate"`
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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
		var equine Equine
		err = json.Unmarshal([]byte(request.Body), &equine)
		if(err != nil){
       return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Bad Request ER01"}, err
		}
		
    // Insert the Equine object into the collection.
    _, err = collection.InsertOne(ctx, equine)
    if err != nil {
        fmt.Printf("Error inserting document: %v\n", err)
        return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Internal Server Error"}, err
    }

		headers := map[string]string{
			"Content-Type":  "application/json",
			"Access-Control-Allow-Origin":      "*", 
			"Access-Control-Allow-Headers":     "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
			"Access-Control-Allow-Methods":     "OPTIONS,POST,GET", 
			"Access-Control-Allow-Credentials": "true",
		}

    return events.APIGatewayProxyResponse{StatusCode: 200, Body: "Equine document inserted successfully", Headers:headers}, nil
}

func main() {
    lambda.Start(Handler)
}