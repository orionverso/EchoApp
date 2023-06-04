package choice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// Plug-in with other constructs
type ChoiceStorageProps struct {
	Storage_solution *string
	Destination      *string
	Granteable       awsiam.IGrantable
}

type choiceStorage struct {
	constructs.Construct
	stg  awsssm.StringParameter
	dest awsssm.StringParameter
	//lo que quieres compartir con otros constructs
}

type ChoiceStorage interface {
	constructs.Construct
	GrantRead(awsiam.IGrantable)
	//define geter() para compartir los objetos
	//ya que new construct devolvera una interface
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

func (ch choiceStorage) GrantRead(gt awsiam.IGrantable) {
	ch.stg.GrantRead(gt)
	ch.dest.GrantRead(gt)
}
