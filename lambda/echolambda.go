package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type receiver interface {
	Write(string) error
}

type s3receiver struct {
	client *s3.Client
}

func (s3rv s3receiver) Write(st string) error {
	body := []byte(st)
	_, err := s3rv.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("DESTINATION")),
		Key:    aws.String("Example.txt"),
		Body:   ioutil.NopCloser(bytes.NewReader(body)),
	})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

type dynamodreceiver struct {
	client *dynamodb.Client
}

func (dbrv dynamodreceiver) Write(st string) error {
	_, err := dbrv.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("DESTINATION")),
		Item: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: st},
		},
	})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

var clientReceiver receiver
var ctx context.Context

func handler(ctx context.Context, ev events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	err := clientReceiver.Write(ev.Body)
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		Body: "Thank you for testing",
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

	if os.Getenv("STORAGE_SOLUTION") == "DYNAMODB" {
		clientReceiver = dynamodreceiver{client: dynamodb.NewFromConfig(cfg)}
	}

	if os.Getenv("STORAGE_SOLUTION") == "S3" {

		clientReceiver = s3receiver{client: s3.NewFromConfig(cfg)}
	}
}
