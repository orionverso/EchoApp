package receiver

import (
	"bytes"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

//Receiver implement writer methods

type Receiver interface {
	Write(context.Context, string) error
}

type s3receiver struct {
	client      *s3.Client
	destination *string
}

func (s3rv s3receiver) Write(ctx context.Context, st string) error {
	body := []byte(st)
	_, err := s3rv.client.PutObject(ctx, &s3.PutObjectInput{
		ContentType: aws.String("application/json"),
		Bucket:      s3rv.destination,
		Key:         aws.String(randstr(10)),
		Body:        bytes.NewReader(body),
	})
	if err != nil {

		log.Println(err)
		return err
	}
	log.Println("data delivered")
	return nil
}

type dynamodbreceiver struct {
	client      *dynamodb.Client
	destination *string
}

func (dbrv dynamodbreceiver) Write(ctx context.Context, st string) error {
	_, err := dbrv.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: dbrv.destination,
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

func GetReceiver(ctx context.Context, cfg aws.Config) (Receiver, error) {
	ssmclient := ssm.NewFromConfig(cfg)
	clientout, err := ssmclient.GetParameter(ctx, &ssm.GetParameterInput{
		Name: aws.String("STORAGE_SOLUTION"),
	})

	if err != nil {
		log.Panicln("The writers don't know to write")
	}

	destout, err := ssmclient.GetParameter(ctx, &ssm.GetParameterInput{
		Name: clientout.Parameter.Value,
	})

	if err != nil {
		log.Panicln("The destination is empty")
	}

	clientType := clientout.Parameter.Value
	dest := destout.Parameter.Value

	if aws.ToString(clientType) == "DYNAMODB" {
		return dynamodbreceiver{
			client:      dynamodb.NewFromConfig(cfg),
			destination: dest,
		}, nil
	}

	if aws.ToString(clientType) == "S3" {
		return s3receiver{
			client:      s3.NewFromConfig(cfg),
			destination: dest,
		}, nil
	}
	return nil, err
}
