package component

import (
	"writer_storage_app/environment"
	"writer_storage_app/storage"
	"writer_storage_app/writer"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ApiLambdaS3Ids struct {
	ApiLambdaS3_Id     string
	WriterApiLambda_Id string
	S3Storage_Id       string
}

type ApiLambdaS3Props struct {
	StackProps           awscdk.StackProps
	WriterApiLambdaProps writer.WriterApiLambdaProps
	S3StorageProps       storage.S3StorageProps
	//Identifiers
	ApiLambdaS3Ids
}

type apiLambdaS3 struct {
	awscdk.Stack
	apiLambda writer.WriterApiLambda
	s3Storage storage.S3Storage
}

func (ap apiLambdaS3) ApiLambda() writer.WriterApiLambda {
	return ap.apiLambda
}

func (ap apiLambdaS3) S3Storage() storage.S3Storage {
	return ap.s3Storage
}

type ApiLambdaS3 interface {
	awscdk.Stack
	ApiLambda() writer.WriterApiLambda
	S3Storage() storage.S3Storage
}

func NewApiLambdaS3(scope constructs.Construct, id *string, props *ApiLambdaS3Props) ApiLambdaS3 {

	var sprops ApiLambdaS3Props = ApiLambdaS3Props_DEFAULT
	var sid ApiLambdaS3Ids = sprops.ApiLambdaS3Ids

	if props != nil {
		sprops = *props
		sid = sprops.ApiLambdaS3Ids
	}

	if id != nil {
		sid.ApiLambdaS3_Id = *id
	}

	stack := awscdk.NewStack(scope, jsii.String(sid.ApiLambdaS3_Id), &sprops.StackProps)

	wr := writer.NewWriterApiLambda(stack, jsii.String(sid.WriterApiLambda_Id), &sprops.WriterApiLambdaProps)

	sprops.S3StorageProps.RoleWriter = wr.Function().Role()

	s3 := storage.NewS3Storage(stack, jsii.String(sid.S3Storage_Id), &sprops.S3StorageProps)
	return apiLambdaS3{stack, wr, s3}
}

// CONFIGURATIONS
var ApiLambdaS3Props_DEFAULT ApiLambdaS3Props = ApiLambdaS3Props{
	StackProps:           environment.StackProps_DEFAULT,
	WriterApiLambdaProps: writer.WriterApiLambdaProps_DEFAULT,
	S3StorageProps:       storage.S3StorageProps_DEFAULT,
	ApiLambdaS3Ids: ApiLambdaS3Ids{
		ApiLambdaS3_Id:     "EchoApp-Implementation-Two-default",
		WriterApiLambda_Id: "WriterApiLambda-Component-default",
		S3Storage_Id:       "RecieveIn-S3Storage-Component-default",
	},
}

var ApiLambdaS3Props_DEV ApiLambdaS3Props = ApiLambdaS3Props{
	StackProps:           environment.StackProps_DEV,
	WriterApiLambdaProps: writer.WriterApiLambdaProps_DEV,
	S3StorageProps:       storage.S3StorageProps_DEV,
	ApiLambdaS3Ids: ApiLambdaS3Ids{
		ApiLambdaS3_Id:     "EchoApp-Implementation-Two-dev",
		WriterApiLambda_Id: "WriterApiLambda-Component-dev",
		S3Storage_Id:       "RecieveIn-S3Storage-Component-dev",
	},
}

var ApiLambdaS3Props_PROD ApiLambdaS3Props = ApiLambdaS3Props{
	StackProps:           environment.StackProps_PROD,
	WriterApiLambdaProps: writer.WriterApiLambdaProps_PROD,
	S3StorageProps:       storage.S3StorageProps_PROD,
	ApiLambdaS3Ids: ApiLambdaS3Ids{
		ApiLambdaS3_Id:     "EchoApp-Implementation-Two-prod",
		WriterApiLambda_Id: "WriterApiLambda-Component-prod",
		S3Storage_Id:       "RecieveIn-S3Storage-Component-prod",
	},
}
