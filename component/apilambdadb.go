package component

import (
	"writer_storage_app/environment"
	"writer_storage_app/storage"
	"writer_storage_app/writer"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ApiLambdaDynamoDbIds struct {
	ApiLambdaDynamoDb_Id string
	WriterApiLambda_Id   string
	DynamoDbStorage_Id   string
}

type ApiLambdaDynamoDbProps struct {
	StackProps           awscdk.StackProps
	WriterApiLambdaProps writer.WriterApiLambdaProps
	DynamoDbStorageProps storage.DynamoDbStorageProps
	//Identifiers
	ApiLambdaDynamoDbIds
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

	var sprops ApiLambdaDynamoDbProps = ApiLambdaDynamoDbProps_DEFAULT
	var sid ApiLambdaDynamoDbIds = sprops.ApiLambdaDynamoDbIds

	if props != nil {
		sprops = *props
		sid = sprops.ApiLambdaDynamoDbIds
	}

	if id != nil {
		sid.ApiLambdaDynamoDb_Id = *id
	}

	stack := awscdk.NewStack(scope, jsii.String(sid.ApiLambdaDynamoDb_Id), &sprops.StackProps)

	wr := writer.NewWriterApiLambda(stack, jsii.String(sid.WriterApiLambda_Id), &sprops.WriterApiLambdaProps)

	sprops.DynamoDbStorageProps.RoleWriter = wr.Function().Role()

	st := storage.NewDynamoDbstorage(stack, jsii.String(sid.DynamoDbStorage_Id), &sprops.DynamoDbStorageProps)

	return apiLambdaDynamoDb{stack, wr, st}
}

// CONFIGURATIONS
var ApiLambdaDynamoDbProps_DEFAULT ApiLambdaDynamoDbProps = ApiLambdaDynamoDbProps{
	StackProps:           environment.StackProps_DEFAULT,
	WriterApiLambdaProps: writer.WriterApiLambdaProps_DEFAULT,
	DynamoDbStorageProps: storage.DynamoDbStorageProps_DEFAULT,
	ApiLambdaDynamoDbIds: ApiLambdaDynamoDbIds{
		ApiLambdaDynamoDb_Id: "EchoApp-Implementation-One-default",
		WriterApiLambda_Id:   "WriterApiLambda-Component-default",
		DynamoDbStorage_Id:   "RecieveIn-DynamoStorage-Component-default",
	},
}

var ApiLambdaDynamoDbProps_DEV ApiLambdaDynamoDbProps = ApiLambdaDynamoDbProps{
	StackProps:           environment.StackProps_DEV,
	WriterApiLambdaProps: writer.WriterApiLambdaProps_DEV,
	DynamoDbStorageProps: storage.DynamoDbStorageProps_DEV,
	ApiLambdaDynamoDbIds: ApiLambdaDynamoDbIds{
		ApiLambdaDynamoDb_Id: "EchoApp-Implementation-One-dev",
		WriterApiLambda_Id:   "WriterApiLambda-Component-dev",
		DynamoDbStorage_Id:   "RecieveIn-DynamoStorage-Component-dev",
	},
}

var ApiLambdaDynamoDbProps_PROD ApiLambdaDynamoDbProps = ApiLambdaDynamoDbProps{
	StackProps:           environment.StackProps_PROD,
	WriterApiLambdaProps: writer.WriterApiLambdaProps_PROD,
	DynamoDbStorageProps: storage.DynamoDbStorageProps_PROD,
	ApiLambdaDynamoDbIds: ApiLambdaDynamoDbIds{
		ApiLambdaDynamoDb_Id: "EchoApp-Implementation-One-prod",
		WriterApiLambda_Id:   "WriterApiLambda-Component-prod",
		DynamoDbStorage_Id:   "RecieveIn-DynamoStorage-Component-prod",
	},
}
