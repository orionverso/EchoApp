package choice

import (
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

//Thanks to choice the writers know where they must to write

type ChoiceStorageProps struct {
	DestinationStringParameterProps awsssm.StringParameterProps
	StorageStringParameterProps     awsssm.StringParameterProps
	//import from interaction with other constructs
	Storage_solution *string
	Destination      *string
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
	this := constructs.NewConstruct(scope, id)

	var sprops ChoiceStorageProps = *props

	sprops.StorageStringParameterProps.StringValue = props.Storage_solution

	stg := awsssm.NewStringParameter(this, jsii.String("STORAGE_SOLUTION_PARAMETER"), &sprops.StorageStringParameterProps)

	sprops.DestinationStringParameterProps.ParameterName = props.Storage_solution
	sprops.DestinationStringParameterProps.StringValue = props.Destination

	dest := awsssm.NewStringParameter(this, jsii.String("DESTINATION_PARAMETER"), &sprops.DestinationStringParameterProps)

	return choiceStorage{this, stg, dest}
}

var ChoiceStorageProps_DEV ChoiceStorageProps = ChoiceStorageProps{
	DestinationStringParameterProps: awsssm.StringParameterProps{},
	StorageStringParameterProps: awsssm.StringParameterProps{
		ParameterName: jsii.String("STORAGE_SOLUTION"),
	},
}

var ChoiceStorageProps_PROD ChoiceStorageProps = ChoiceStorageProps{
	DestinationStringParameterProps: awsssm.StringParameterProps{},
	StorageStringParameterProps: awsssm.StringParameterProps{
		ParameterName: jsii.String("STORAGE_SOLUTION"),
	},
}
