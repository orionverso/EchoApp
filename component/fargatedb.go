package component

import (
	"writer_storage_app/environment"
	"writer_storage_app/storage"
	"writer_storage_app/writer"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FargateDynamoDbProps struct {
	StackProps           awscdk.StackProps
	WriterFargateProps   writer.WriterFargateProps
	DynamoDbStorageProps storage.DynamoDbStorageProps
}

type fargateDynamoDb struct {
	awscdk.Stack
	fargateService writer.WriterFargate
	dynamoStorage  storage.DynamoDbStorage
}

func (fg fargateDynamoDb) Fargate() writer.WriterFargate {
	return fg.fargateService
}

func (fg fargateDynamoDb) DynamoStorage() storage.DynamoDbStorage {
	return fg.dynamoStorage
}

type FargateDynamoDb interface {
	awscdk.Stack
	Fargate() writer.WriterFargate
	DynamoStorage() storage.DynamoDbStorage
}

func NewFargateDynamoDb(scope constructs.Construct, id *string, props *FargateDynamoDbProps) FargateDynamoDb {
	var sprops FargateDynamoDbProps = FargateDynamoDbProps_DEV
	if props != nil {
		sprops = *props
	}
	stack := awscdk.NewStack(scope, id, &sprops.StackProps)

	wr := writer.NewWriterFargate(stack, jsii.String("TaskWriter"), &sprops.WriterFargateProps)

	sprops.DynamoDbStorageProps.RoleWriter = wr.FargateService().TaskDefinition().ExecutionRole()

	st := storage.NewDynamoDbstorage(stack, jsii.String("DynamoDbStorage"), &sprops.DynamoDbStorageProps)

	return fargateDynamoDb{stack, wr, st}
}

// CONFIGURATIONS
var FargateDynamoDbProps_DEV FargateDynamoDbProps = FargateDynamoDbProps{
	StackProps:           environment.StackProps_DEV,
	WriterFargateProps:   writer.WriterFargateProps_DEV,
	DynamoDbStorageProps: storage.DynamoDbStorageProps_DEV,
}

var FargateDynamoDbProps_PROD FargateDynamoDbProps = FargateDynamoDbProps{
	StackProps:           environment.StackProps_PROD,
	WriterFargateProps:   writer.WriterFargateProps_PROD,
	DynamoDbStorageProps: storage.DynamoDbStorageProps_PROD,
}
