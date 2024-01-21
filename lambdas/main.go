package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type App struct {
	id string
}


func newApp(id string) *App {
	return &App{
		id:id,
	}
}

func (app *App) Handler(request events.APIGatewayProxyRequest ) (events.APIGatewayProxyResponse, error) {
	responseBody := map[string]string{
		"message":"Hi you have hit the route",
	}

	responseJSON, err := json.Marshal(responseBody)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body:       `{"error":"internal server error code ER01"}`,
			IsBase64Encoded: false,
		}, err
	}

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Access-Control-Allow-Origin":      "*", 
		"Access-Control-Allow-Headers":     "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
		"Access-Control-Allow-Methods":     "OPTIONS,POST,GET", 
		"Access-Control-Allow-Credentials": "true",
	}
	response := events.APIGatewayProxyResponse{
		Body: string(responseJSON),
		StatusCode: http.StatusOK,
		Headers: headers,
	}

	return response, nil
}

func main(){
	id := "some random id"
	app := newApp(id)

	lambda.Start(app.Handler)
}