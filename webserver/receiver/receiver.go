package receiver

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type getReceiverProps struct {
	stgGetParameterInput ssm.GetParameterInput
	dstGetParameterInput ssm.GetParameterInput
}

type Receiver interface {
	Write(ctx context.Context, st string) error
}

func GetReceiver(ctx context.Context, cfg aws.Config) (Receiver, error) {
	var sprops getReceiverProps = getReceiverProps_DEFAULT

	ssmclient := ssm.NewFromConfig(cfg)
	stg, err := ssmclient.GetParameter(ctx, &sprops.stgGetParameterInput)

	if err != nil {
		log.Println(err)
	}

	storage := aws.ToString(stg.Parameter.Value)
	log.Println("Storage:", storage)

	sprops.dstGetParameterInput.Name = stg.Parameter.Value
	dst, err := ssmclient.GetParameter(ctx, &sprops.dstGetParameterInput)

	if err != nil {
		log.Println(err)
	}

	destination := aws.ToString(dst.Parameter.Value)
	log.Println("Destination:", destination)

	if storage == "DYNAMODB" {
		return dynamoDbReceiver{
			*dynamodb.NewFromConfig(cfg),
			destination,
		}, nil
	}

	if storage == "S3" {
		return s3Receiver{
			*s3.NewFromConfig(cfg),
			destination,
		}, nil
	}

	return nil, err //nil pointer desreference
}

// CONFIGURATIONS

var getReceiverProps_DEFAULT getReceiverProps = getReceiverProps{
	stgGetParameterInput: ssm.GetParameterInput{
		Name: aws.String("STORAGE_SOLUTION"),
	},
	dstGetParameterInput: ssm.GetParameterInput{},
}
