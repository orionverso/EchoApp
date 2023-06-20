package choice

import (
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// Thanks to choice the writers know where they must to write
type ChoiceStorageIds struct {
	ChoiceStorage_Id     string
	StorageParameter     string
	DestinationParameter string
}

type ChoiceStorageProps struct {
	DestinationStringParameterProps awsssm.StringParameterProps
	StorageStringParameterProps     awsssm.StringParameterProps
	//import from interaction with other constructs
	Storage_solution *string
	Destination      *string
	//Identifiers
	ChoiceStorageIds
}

type choiceStorage struct {
	constructs.Construct
	storage     awsssm.StringParameter
	destination awsssm.StringParameter
}

func (ch choiceStorage) Storage() awsssm.StringParameter {
	return ch.storage
}

func (ch choiceStorage) Destination() awsssm.StringParameter {
	return ch.destination
}

type ChoiceStorage interface {
	constructs.Construct
	Storage() awsssm.StringParameter
	Destination() awsssm.StringParameter
}

func NewChoiceStorage(scope constructs.Construct, id *string, props *ChoiceStorageProps) ChoiceStorage {

	if props == nil {
		log.Panicln("You must define Storage and Destination\n. Writers won't know where to send data")
	}

	var sprops ChoiceStorageProps = *props
	var sid ChoiceStorageIds = sprops.ChoiceStorageIds

	if id != nil {
		sid.ChoiceStorage_Id = *id
	}

	this := constructs.NewConstruct(scope, jsii.String(sid.ChoiceStorage_Id))

	sprops.StorageStringParameterProps.StringValue = props.Storage_solution

	stg := awsssm.NewStringParameter(this, jsii.String(sid.StorageParameter), &sprops.StorageStringParameterProps)

	sprops.DestinationStringParameterProps.ParameterName = sprops.Storage_solution
	sprops.DestinationStringParameterProps.StringValue = sprops.Destination

	dest := awsssm.NewStringParameter(this, jsii.String(sid.DestinationParameter), &sprops.DestinationStringParameterProps)

	return choiceStorage{this, stg, dest}
}

// CONFIGURATIONS
var ChoiceStorageProps_DEFAULT ChoiceStorageProps = ChoiceStorageProps{
	DestinationStringParameterProps: awsssm.StringParameterProps{},
	StorageStringParameterProps: awsssm.StringParameterProps{
		ParameterName: jsii.String("STORAGE_SOLUTION"),
	},
	ChoiceStorageIds: ChoiceStorageIds{
		ChoiceStorage_Id:     "StorageChoice-FromUnknown-default",
		StorageParameter:     "STORAGE_SOLUTION_PARAMETER-default",
		DestinationParameter: "DESTINATION_PARAMETER-default",
	},
}

var ChoiceStorageProps_DEV ChoiceStorageProps = ChoiceStorageProps{
	DestinationStringParameterProps: awsssm.StringParameterProps{},
	StorageStringParameterProps: awsssm.StringParameterProps{
		ParameterName: jsii.String("STORAGE_SOLUTION"),
	},
	ChoiceStorageIds: ChoiceStorageIds{
		ChoiceStorage_Id:     "StorageChoice-FromUnknown-dev",
		StorageParameter:     "STORAGE_SOLUTION_PARAMETER-dev",
		DestinationParameter: "DESTINATION_PARAMETER-default-dev",
	},
}

var ChoiceStorageProps_PROD ChoiceStorageProps = ChoiceStorageProps{
	DestinationStringParameterProps: awsssm.StringParameterProps{},
	StorageStringParameterProps: awsssm.StringParameterProps{
		ParameterName: jsii.String("STORAGE_SOLUTION"),
	},
	ChoiceStorageIds: ChoiceStorageIds{
		ChoiceStorage_Id:     "StorageChoice-FromUnknown-prod",
		StorageParameter:     "STORAGE_SOLUTION_PARAMETER-prod",
		DestinationParameter: "DESTINATION_PARAMETER-default-prod",
	},
}
