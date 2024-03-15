package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Input struct {
	Name string
}

func main() {
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var input Input
		err := json.Unmarshal([]byte(request.Body), &input)
		if err != nil {
			return events.APIGatewayProxyResponse{Body: "error", StatusCode: 400}, nil
		}

		return events.APIGatewayProxyResponse{Body: fmt.Sprintf("Hello %s", input.Name), StatusCode: 200}, nil
	})
}
