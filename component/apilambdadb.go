package component

import (
	"writer_storage_app/storage"
	"writer_storage_app/writer"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ApiLambdaDynamoDbProps struct {
	awscdk.StackProps
}

type apiLambdaDynamoDb struct {
	awscdk.Stack
	apiLambda     writer.WriterApiLambda
	dynamoStorage storage.DynamoDbStorage
}

func (ap apiLambdaDynamoDb) ApiLambda() writer.WriterApiLambda {
	return ap.apiLambda
}

func (ap apiLambdaDynamoDb) DynamoStorage() storage.DynamoDbStorage {
	return ap.dynamoStorage
}

type ApiLambdaDynamoDb interface {
	awscdk.Stack
	ApiLambda() writer.WriterApiLambda
	DynamoStorage() storage.DynamoDbStorage
}

func NewApiLambdaDynamoDb(scope constructs.Construct, id *string, props *ApiLambdaDynamoDbProps) ApiLambdaDynamoDb {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, id, &sprops)

	wr := writer.NewWriterApiLambda(stack, jsii.String("LambdaApiWriter"), &writer.WriterApiLambdaProps{})

	st := storage.NewDynamoDbstorage(stack, jsii.String("DynamoDbStorage"), &storage.DynamoDbstorageProps{
		RoleWriter: wr.Function().Role(),
	})

	return apiLambdaDynamoDb{stack, wr, st}
}
