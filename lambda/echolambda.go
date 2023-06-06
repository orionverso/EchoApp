package main

import (
	"context"
	"log"

	"writer_storage_app/storage/receiver"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
)

var clientReceiver receiver.Receiver
var ctx context.Context

func handler(ctx context.Context, ev events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	err := clientReceiver.Write(ctx, ev.Body)
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		IsBase64Encoded:   false,
		StatusCode:        200,
		Headers:           map[string]string{},
		MultiValueHeaders: map[string][]string{},
		Body:              "Thank you for take a look. I am from Lambda.See you",
	}, nil
}

func main() {
	lambda.Start(handler)
}

func init() {
	ctx = context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		// Handle error
	}
	clientReceiver, err = receiver.GetReceiver(ctx, cfg)
}
