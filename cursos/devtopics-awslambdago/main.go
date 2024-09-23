package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
}

type MyEv struct {
	Name string `json:"name"`
}

type DefaultResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func MyHandler(ctx context.Context, event events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	var res *events.APIGatewayProxyResponse

	logger.Info("received event", zap.Any("method", event.HTTPMethod), zap.Any("path", event.Path), zap.Any("body", event.Body))

	if event.Path == "/hello" {

		body, _ := json.Marshal(&DefaultResponse{
			Status:  string(http.StatusOK),
			Message: "Hello World",
		})

		res = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       string(body),
		}
	} else {
		res = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       "Outra coisa qualquer",
		}
	}

	return res, nil
}

func main() {
	lambda.Start(MyHandler)
}
