package storage

import (
	"log"
	"writer_storage_app/storage/choice"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// Plug-in with other constructs
type DynamoDbstorageProps struct {
	PlugWriter any
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

	ch := choice.NewChoiceStorage(this, jsii.String("StorageChoice"), &choice.ChoiceStorageProps{
		Storage_solution: jsii.String("DYNAMODB"),
		Destination:      table.TableName(),
	})
	//TYPE ASSERTION IN ORDER TO KNOW POSSIBLE WRITER AND GRANT PERMISION
	plfn, ok := props.PlugWriter.(awslambda.Function)

	if ok {
		table.GrantWriteData(plfn)
		ch.GrantRead(plfn)
	}

	plserv, ok := props.PlugWriter.(awsecspatterns.ApplicationLoadBalancedFargateService)

	if ok {
		table.GrantWriteData(plserv.TaskDefinition().TaskRole())
		ch.GrantRead(plserv.TaskDefinition().TaskRole())
	}

	/*
		type assertion new writer here
	*/

	if !ok {
		log.Panicln("You must define a writer for storage solution")
	}

	return dynamoDbstorage{this}
}
