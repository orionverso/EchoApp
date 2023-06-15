package component

import (
	"writer_storage_app/storage"
	"writer_storage_app/writer"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ApiLambdaS3Props struct {
	awscdk.StackProps
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
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, id, &sprops)

	wr := writer.NewWriterApiLambda(stack, jsii.String("LambdaApiWriter"), &writer.WriterApiLambdaProps{})

	s3 := storage.NewS3Storage(stack, jsii.String("S3Storage"), &storage.S3StorageProps{
		RoleWriter: wr.Function().Role(),
	})

	return apiLambdaS3{stack, wr, s3}
}
