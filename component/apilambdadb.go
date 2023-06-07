package component

import (
	"writer_storage_app/storage"
	"writer_storage_app/writer"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type WriterStorageAppStackApiLambdaDBProps struct {
	awscdk.StackProps
}

func NewWriterStorageAppStackApiLambdaDB(scope constructs.Construct, id *string, props *WriterStorageAppStackApiLambdaDBProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, id, &sprops)

	wr := writer.NewWriterApiLambda(stack, jsii.String("LambdaApiWriter"), &writer.WriterApiLambdaProps{})

	storage.NewDynamoDbstorage(stack, jsii.String("DynamoDbStorage"), &storage.DynamoDbstorageProps{
		PlugGranteableWriter: wr.PlugGranteableFunc(),
	})

	return stack
}
