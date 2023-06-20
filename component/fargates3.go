package component

import (
	"writer_storage_app/environment"
	"writer_storage_app/storage"
	"writer_storage_app/writer"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FargateS3Ids struct {
	FargateS3_Id     string
	WriterFargate_Id string
	S3Storage_Id     string
}

type FargateS3Props struct {
	StackProps         awscdk.StackProps
	WriterFargateProps writer.WriterFargateProps
	S3StorageProps     storage.S3StorageProps
	//Identifiers
	FargateS3Ids
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

	var sprops FargateS3Props = FargateS3Props_DEFAULT
	var sid FargateS3Ids = sprops.FargateS3Ids

	if props != nil {
		sprops = *props
		sid = sprops.FargateS3Ids
	}

	if id != nil {
		sid.FargateS3_Id = *id
	}

	stack := awscdk.NewStack(scope, jsii.String(sid.FargateS3_Id), &sprops.StackProps)

	wr := writer.NewWriterFargate(stack, jsii.String(sid.FargateS3_Id), &sprops.WriterFargateProps)

	sprops.S3StorageProps.RoleWriter = wr.FargateService().TaskDefinition().TaskRole()

	st := storage.NewS3Storage(stack, jsii.String(sid.S3Storage_Id), &sprops.S3StorageProps)

	return fargateS3{stack, wr, st}
}

// CONFIGURATIONS
var FargateS3Props_DEFAULT FargateS3Props = FargateS3Props{
	StackProps:         environment.StackProps_DEFAULT,
	WriterFargateProps: writer.WriterFargateProps_DEFAULT,
	S3StorageProps:     storage.S3StorageProps_DEFAULT,
	FargateS3Ids: FargateS3Ids{
		FargateS3_Id:     "EchoApp-EchoApp-Implementation-Four-default",
		WriterFargate_Id: "WriterFargate-component-default",
		S3Storage_Id:     "RecieveIn-S3Storage-Component-default",
	},
}

var FargateS3Props_DEV FargateS3Props = FargateS3Props{
	StackProps:         environment.StackProps_DEV,
	WriterFargateProps: writer.WriterFargateProps_DEV,
	S3StorageProps:     storage.S3StorageProps_DEV,
	FargateS3Ids: FargateS3Ids{
		FargateS3_Id:     "EchoApp-EchoApp-Implementation-Four-dev",
		WriterFargate_Id: "WriterFargate-component-dev",
		S3Storage_Id:     "RecieveIn-S3Storage-Component-dev",
	},
}

var FargateS3Props_PROD FargateS3Props = FargateS3Props{
	StackProps:         environment.StackProps_PROD,
	WriterFargateProps: writer.WriterFargateProps_PROD,
	S3StorageProps:     storage.S3StorageProps_PROD,
	FargateS3Ids: FargateS3Ids{
		FargateS3_Id:     "EchoApp-EchoApp-Implementation-Four-prod",
		WriterFargate_Id: "WriterFargate-component-prod",
		S3Storage_Id:     "RecieveIn-S3Storage-Component-prod",
	},
}
