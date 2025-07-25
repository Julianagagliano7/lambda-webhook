package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received request:", request.Body)

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       request.Body,
	}

	return response, nil
}
