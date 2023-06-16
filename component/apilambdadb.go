package component

import (
	"writer_storage_app/enviroment"
	"writer_storage_app/storage"
	"writer_storage_app/writer"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ApiLambdaDynamoDbProps struct {
	StackProps           awscdk.StackProps
	WriterApiLambdaProps writer.WriterApiLambdaProps
	DynamoDbStorageProps storage.DynamoDbStorageProps
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
	var sprops ApiLambdaDynamoDbProps = ApiLambdaDynamoDbProps_DEV

	if props != nil {
		sprops = *props
	}
	stack := awscdk.NewStack(scope, id, &sprops.StackProps)

	wr := writer.NewWriterApiLambda(stack, jsii.String("LambdaApiWriter"), &sprops.WriterApiLambdaProps)

	sprops.DynamoDbStorageProps.RoleWriter = wr.Function().Role()

	st := storage.NewDynamoDbstorage(stack, jsii.String("DynamoDbStorage"), &sprops.DynamoDbStorageProps)

	return apiLambdaDynamoDb{stack, wr, st}
}

// CONFIGURATIONS
var ApiLambdaDynamoDbProps_DEV ApiLambdaDynamoDbProps = ApiLambdaDynamoDbProps{
	StackProps:           enviroment.StackProps_DEV,
	WriterApiLambdaProps: writer.WriterApiLambdaProps_DEV,
	DynamoDbStorageProps: storage.DynamoDbStorageProps_DEV,
}

var ApiLambdaDynamoDbProps_PROD ApiLambdaDynamoDbProps = ApiLambdaDynamoDbProps{
	StackProps:           enviroment.StackProps_PROD,
	WriterApiLambdaProps: writer.WriterApiLambdaProps_PROD,
	DynamoDbStorageProps: storage.DynamoDbStorageProps_PROD,
}
