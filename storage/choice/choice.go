package choice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

//Thanks to choice the writers know where they must to write

type ChoiceStorageProps struct {
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
	//implement construct
	this := constructs.NewConstruct(scope, id)

	stg := awsssm.NewStringParameter(this, jsii.String("STORAGE_SOLUTION_PARAMETER"), &awsssm.StringParameterProps{
		ParameterName: jsii.String("STORAGE_SOLUTION"),
		StringValue:   props.Storage_solution,
	})

	dest := awsssm.NewStringParameter(this, jsii.String("DESTINATION_PARAMETER"), &awsssm.StringParameterProps{
		ParameterName: props.Storage_solution,
		StringValue:   props.Destination,
	})

	return choiceStorage{this, stg, dest}
}
