package component

import (
	"writer_storage_app/environment"
	"writer_storage_app/storage"
	"writer_storage_app/writer"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ApiLambdaS3Props struct {
	StackProps           awscdk.StackProps
	WriterApiLambdaProps writer.WriterApiLambdaProps
	S3StorageProps       storage.S3StorageProps
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
	var sprops ApiLambdaS3Props = ApiLambdaS3Props_DEV

	if props != nil {
		sprops = *props
	}

	stack := awscdk.NewStack(scope, id, &sprops.StackProps)

	wr := writer.NewWriterApiLambda(stack, jsii.String("LambdaApiWriter"), &sprops.WriterApiLambdaProps)

	sprops.S3StorageProps.RoleWriter = wr.Function().Role()

	s3 := storage.NewS3Storage(stack, jsii.String("S3Storage"), &sprops.S3StorageProps)
	return apiLambdaS3{stack, wr, s3}
}

// CONFIGURATIONS
var ApiLambdaS3Props_DEV ApiLambdaS3Props = ApiLambdaS3Props{
	StackProps:           environment.StackProps_DEV,
	WriterApiLambdaProps: writer.WriterApiLambdaProps_DEV,
	S3StorageProps:       storage.S3StorageProps_DEV,
}

var ApiLambdaS3Props_PROD ApiLambdaS3Props = ApiLambdaS3Props{
	StackProps:           environment.StackProps_PROD,
	WriterApiLambdaProps: writer.WriterApiLambdaProps_PROD,
	S3StorageProps:       storage.S3StorageProps_PROD,
}
