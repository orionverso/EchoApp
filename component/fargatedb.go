package component

import (
	"writer_storage_app/storage"
	"writer_storage_app/writer"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FargateDynamoDbProps struct {
	awscdk.StackProps
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
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, id, &sprops)

	wr := writer.NewWriterFargate(stack, jsii.String("TaskWriter"), &writer.WriterFargateProps{})

	st := storage.NewDynamoDbstorage(stack, jsii.String("DynamoDbStorage"), &storage.DynamoDbstorageProps{
		RoleWriter: wr.FargateService().TaskDefinition().TaskRole(),
	})
	return fargateDynamoDb{stack, wr, st}
}
