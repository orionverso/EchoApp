package component

import (
	"writer_storage_app/environment"
	"writer_storage_app/storage"
	"writer_storage_app/writer"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FargateDynamoDbIds struct {
	FargateDynamoDb_Id string
	WriterFargate_Id   string
	DynamoStorage_Id   string
}

type FargateDynamoDbProps struct {
	StackProps           awscdk.StackProps
	WriterFargateProps   writer.WriterFargateProps
	DynamoDbStorageProps storage.DynamoDbStorageProps
	//Identifiers
	FargateDynamoDbIds
}

type fargateDynamoDb struct {
	awscdk.Stack
	fargateService writer.WriterFargate
	dynamoStorage  storage.DynamoDbStorage
}

func (fg fargateDynamoDb) Fargate() writer.WriterFargate {
	return fg.fargateService
}

func (fg fargateDynamoDb) DynamoStorage() storage.DynamoDbStorage {
	return fg.dynamoStorage
}

type FargateDynamoDb interface {
	awscdk.Stack
	Fargate() writer.WriterFargate
	DynamoStorage() storage.DynamoDbStorage
}

func NewFargateDynamoDb(scope constructs.Construct, id *string, props *FargateDynamoDbProps) FargateDynamoDb {

	var sprops FargateDynamoDbProps = FargateDynamoDbProps_DEFAULT
	var sid FargateDynamoDbIds = sprops.FargateDynamoDbIds

	if props != nil {
		sprops = *props
		sid = sprops.FargateDynamoDbIds
	}

	if id != nil {
		sid.FargateDynamoDb_Id = *id
	}

	stack := awscdk.NewStack(scope, jsii.String(sid.FargateDynamoDb_Id), &sprops.StackProps)

	wr := writer.NewWriterFargate(stack, jsii.String(sid.WriterFargate_Id), &sprops.WriterFargateProps)

	sprops.DynamoDbStorageProps.RoleWriter = wr.FargateService().TaskDefinition().TaskRole()

	st := storage.NewDynamoDbstorage(stack, jsii.String(sid.DynamoStorage_Id), &sprops.DynamoDbStorageProps)

	return fargateDynamoDb{stack, wr, st}
}

// CONFIGURATIONS
var FargateDynamoDbProps_DEFAULT FargateDynamoDbProps = FargateDynamoDbProps{
	StackProps:           environment.StackProps_DEFAULT,
	WriterFargateProps:   writer.WriterFargateProps_DEFAULT,
	DynamoDbStorageProps: storage.DynamoDbStorageProps_DEFAULT,
	FargateDynamoDbIds: FargateDynamoDbIds{
		FargateDynamoDb_Id: "EchoApp-Implementation-Three-default",
		WriterFargate_Id:   "WriterFargate-Component-default",
		DynamoStorage_Id:   "RecieveIn-DynamoStorage-Component-default",
	},
}

var FargateDynamoDbProps_DEV FargateDynamoDbProps = FargateDynamoDbProps{
	StackProps:           environment.StackProps_DEV,
	WriterFargateProps:   writer.WriterFargateProps_DEV,
	DynamoDbStorageProps: storage.DynamoDbStorageProps_DEV,
	FargateDynamoDbIds: FargateDynamoDbIds{
		FargateDynamoDb_Id: "EchoApp-Implementation-Three-dev",
		WriterFargate_Id:   "WriterFargate-Component-dev",
		DynamoStorage_Id:   "RecieveIn-DynamoStorage-Component-dev",
	},
}

var FargateDynamoDbProps_PROD FargateDynamoDbProps = FargateDynamoDbProps{
	StackProps:           environment.StackProps_PROD,
	WriterFargateProps:   writer.WriterFargateProps_PROD,
	DynamoDbStorageProps: storage.DynamoDbStorageProps_PROD,
	FargateDynamoDbIds: FargateDynamoDbIds{
		FargateDynamoDb_Id: "EchoApp-Implementation-Three-prod",
		WriterFargate_Id:   "WriterFargate-Component-prod",
		DynamoStorage_Id:   "RecieveIn-DynamoStorage-Component-prod",
	},
}
