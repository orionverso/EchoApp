package main

import (
	"bytes"
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
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
		ContentType: aws.String("application/json"),
		Bucket:      dest,
		Key:         aws.String("Example.json"),
		Body:        bytes.NewReader(body),
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
		TableName: dest,
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
var dest *string

func handler(ctx context.Context, ev events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	err := clientReceiver.Write(ev.Body)
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		IsBase64Encoded:   false,
		StatusCode:        200,
		Headers:           map[string]string{},
		MultiValueHeaders: map[string][]string{},
		Body:              "Thank you for testing",
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

	ssmclient := ssm.NewFromConfig(cfg)

	stgout, err := ssmclient.GetParameter(ctx, &ssm.GetParameterInput{
		Name: aws.String("STORAGE_SOLUTION"),
	})

	destout, err := ssmclient.GetParameter(ctx, &ssm.GetParameterInput{
		Name: stgout.Parameter.Value,
	})

	dest = destout.Parameter.Value

	if err != nil {
		// Handle error
	}

	if aws.ToString(stgout.Parameter.Value) == "DYNAMODB" {
		clientReceiver = dynamodreceiver{client: dynamodb.NewFromConfig(cfg)}
	}

	if aws.ToString(stgout.Parameter.Value) == "S3" {
		clientReceiver = s3receiver{client: s3.NewFromConfig(cfg)}
	}

}
