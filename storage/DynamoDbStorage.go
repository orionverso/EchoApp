package storage

import (
	"writer_storage_app/storage/choice"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// Plug-in with other constructs
type DynamoDbstorageProps struct {
	PlugFunc awslambda.Function
	//insert props from other constructs
}

type dynamoDbstorage struct {
	constructs.Construct
	//insert construct from other resources
}

type DynamoDbstorage interface {
	constructs.Construct
	//insert useful method to Do construct
}

func NewDynamoDbstorage(scope constructs.Construct, id *string, props *DynamoDbstorageProps) DynamoDbstorage {
	//implement construct
	this := constructs.NewConstruct(scope, id)

	table := awsdynamodb.NewTable(this, jsii.String("ReceiveEchoTable"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})
	//Permision to write
	table.GrantWriteData(props.PlugFunc)
	//Add destination for lambda at Run time

	choice.NewChoiceStorage(this, jsii.String("StorageChoice"), &choice.ChoiceStorageProps{
		Storage_solution: jsii.String("DYNAMODB"),
		Destination:      table.TableName(),
		Granteable:       props.PlugFunc,
	})

	return dynamoDbstorage{this}
}
