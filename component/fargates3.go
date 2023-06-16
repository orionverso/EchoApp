package component

import (
	"writer_storage_app/environment"
	"writer_storage_app/storage"
	"writer_storage_app/writer"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FargateS3Props struct {
	StackProps         awscdk.StackProps
	WriterFargateProps writer.WriterFargateProps
	S3StorageProps     storage.S3StorageProps
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
	var sprops FargateS3Props = FargateS3Props_DEV
	if props != nil {
		sprops = *props
	}
	stack := awscdk.NewStack(scope, id, &sprops.StackProps)

	wr := writer.NewWriterFargate(stack, jsii.String("TaskWriter"), &sprops.WriterFargateProps)

	sprops.S3StorageProps.RoleWriter = wr.FargateService().TaskDefinition().ExecutionRole()

	st := storage.NewS3Storage(stack, jsii.String("S3Storage"), &sprops.S3StorageProps)

	return fargateS3{stack, wr, st}
}

var FargateS3Props_DEV FargateS3Props = FargateS3Props{
	StackProps:         environment.StackProps_DEV,
	WriterFargateProps: writer.WriterFargateProps_DEV,
	S3StorageProps:     storage.S3StorageProps_DEV,
}

var FargateS3Props_PROD FargateS3Props = FargateS3Props{
	StackProps:         environment.StackProps_PROD,
	WriterFargateProps: writer.WriterFargateProps_PROD,
	S3StorageProps:     storage.S3StorageProps_PROD,
}
