package storage

import (
	"writer_storage_app/storage/choice"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type DynamoDbStorageIds struct {
	DynamoDbStorage_Id string
	Table_Id           string
	Choice_Id          string
}

type DynamoDbStorageProps struct {
	TableProps         awsdynamodb.TableProps
	ChoiceStorageProps choice.ChoiceStorageProps
	//import from interaction with other constructs
	RoleWriter awsiam.IRole
	//Identifiers
	DynamoDbStorageIds
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

	var sprops DynamoDbStorageProps = DynamoDbStorageProps_DEFAULT
	var sid DynamoDbStorageIds = sprops.DynamoDbStorageIds

	if props != nil {
		sprops = *props
		sid = sprops.DynamoDbStorageIds
	}

	if id != nil {
		sid.DynamoDbStorage_Id = *id
	}

	this := constructs.NewConstruct(scope, jsii.String(sid.DynamoDbStorage_Id))

	table := awsdynamodb.NewTable(this, jsii.String(sid.Table_Id), &sprops.TableProps)

	sprops.ChoiceStorageProps.Storage_solution = jsii.String("DYNAMODB")
	sprops.ChoiceStorageProps.Destination = table.TableName()

	ch := choice.NewChoiceStorage(this, jsii.String(sid.Choice_Id), &sprops.ChoiceStorageProps)

	table.GrantWriteData(props.RoleWriter)
	ch.Storage().GrantRead(props.RoleWriter)
	ch.Destination().GrantRead(props.RoleWriter)

	return dynamoDbStorage{this, table, ch}
}

// CONFIGURATIONS
var DynamoDbStorageProps_DEFAULT DynamoDbStorageProps = DynamoDbStorageProps{
	TableProps: awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	},
	ChoiceStorageProps: choice.ChoiceStorageProps_DEFAULT,
	DynamoDbStorageIds: DynamoDbStorageIds{
		DynamoDbStorage_Id: "RecieveIn-DynamoStorage-default",
		Table_Id:           "ReceiveEchoTable-default",
		Choice_Id:          "StorageChoice-FromTable-default",
	},
}

var DynamoDbStorageProps_DEV DynamoDbStorageProps = DynamoDbStorageProps{
	TableProps: awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	},
	ChoiceStorageProps: choice.ChoiceStorageProps_DEV,
	DynamoDbStorageIds: DynamoDbStorageIds{
		DynamoDbStorage_Id: "RecieveIn-DynamoStorage-dev",
		Table_Id:           "ReceiveEchoTable-dev",
		Choice_Id:          "StorageChoice-FromTable-dev",
	},
}

var DynamoDbStorageProps_PROD DynamoDbStorageProps = DynamoDbStorageProps{
	TableProps: awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	},
	ChoiceStorageProps: choice.ChoiceStorageProps_PROD,
	DynamoDbStorageIds: DynamoDbStorageIds{
		DynamoDbStorage_Id: "RecieveIn-DynamoStorage-prod",
		Table_Id:           "ReceiveEchoTable-prod",
		Choice_Id:          "StorageChoice-FromTable-prod",
	},
}
