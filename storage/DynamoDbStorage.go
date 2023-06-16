package storage

import (
	"writer_storage_app/storage/choice"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type DynamoDbStorageProps struct {
	TableProps         awsdynamodb.TableProps
	ChoiceStorageProps choice.ChoiceStorageProps
	//import from interaction with other constructs
	RoleWriter awsiam.IRole
}

type dynamoDbStorage struct {
	constructs.Construct
	table  awsdynamodb.Table
	choice choice.ChoiceStorage
}

func (db dynamoDbStorage) Table() awsdynamodb.Table {
	return db.table
}

func (db dynamoDbStorage) Choice() choice.ChoiceStorage {
	return db.choice
}

type DynamoDbStorage interface {
	constructs.Construct
	Table() awsdynamodb.Table
	Choice() choice.ChoiceStorage
}

func NewDynamoDbstorage(scope constructs.Construct, id *string, props *DynamoDbStorageProps) DynamoDbStorage {
	var sprops DynamoDbStorageProps = DynamoDbstorageProps_DEV
	if props != nil {
		sprops = *props
	}

	this := constructs.NewConstruct(scope, id)

	table := awsdynamodb.NewTable(this, jsii.String("ReceiveEchoTable"), &sprops.TableProps)

	sprops.ChoiceStorageProps.Storage_solution = jsii.String("DYNAMODB")
	sprops.ChoiceStorageProps.Destination = table.TableName()

	ch := choice.NewChoiceStorage(this, jsii.String("StorageChoice"), &sprops.ChoiceStorageProps)

	table.GrantWriteData(props.RoleWriter)
	ch.Storage().GrantRead(props.RoleWriter)
	ch.Destination().GrantRead(props.RoleWriter)

	return dynamoDbStorage{this, table, ch}
}

// CONFIGURATIONS
var DynamoDbStorageProps_DEV DynamoDbStorageProps = DynamoDbStorageProps{
	TableProps: awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	},
	ChoiceStorageProps: choice.ChoiceStorageProps_DEV,
}

var DynamoDbStorageProps_PROD DynamoDbStorageProps = DynamoDbStorageProps{
	TableProps: awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	},
	ChoiceStorageProps: choice.ChoiceStorageProps_PROD,
}
