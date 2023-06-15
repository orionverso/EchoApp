package component

import (
	"writer_storage_app/storage"
	"writer_storage_app/writer"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FargateS3Props struct {
	awscdk.StackProps
}

type fargateS3 struct {
	awscdk.Stack
	fargateService writer.WriterFargate
	s3Storage      storage.S3Storage
}

func (fg fargateS3) Fargate() writer.WriterFargate {
	return fg.fargateService
}

func (fg fargateS3) S3Storage() storage.S3Storage {
	return fg.s3Storage
}

type FargateS3 interface {
	awscdk.Stack
	Fargate() writer.WriterFargate
	S3Storage() storage.S3Storage
}

func NewFargateS3(scope constructs.Construct, id *string, props *FargateS3Props) FargateS3 {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, id, &sprops)

	wr := writer.NewWriterFargate(stack, jsii.String("TaskWriter"), &writer.WriterFargateProps{})

	st := storage.NewS3Storage(stack, jsii.String("S3Storage"), &storage.S3StorageProps{
		RoleWriter: wr.FargateService().TaskDefinition().TaskRole(),
	})

	return fargateS3{stack, wr, st}
}
