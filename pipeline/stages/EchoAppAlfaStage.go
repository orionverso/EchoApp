package stages

import (
	"writer_storage_app/component"
	"writer_storage_app/environment"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type EchoAppAlfaIds struct {
	EchoAppAlfa_Id       string
	ApiLambdaDynamoDb_Id string
}

type EchoAppAlfaProps struct {
	StageProps             awscdk.StageProps
	ApiLambdaDynamoDbProps component.ApiLambdaDynamoDbProps
	//Identifiers
	EchoAppAlfaIds
}

type echoAppAlfa struct {
	awscdk.Stage
	cpt component.ApiLambdaDynamoDb
}

func (ec echoAppAlfa) EchoAppAlfaComponent() component.ApiLambdaDynamoDb {
	return ec.cpt
}

func (ec echoAppAlfa) EchoAppAlfaStage() awscdk.Stage {
	return ec.Stage
}

type EchoAppAlfa interface {
	awscdk.Stage
	EchoAppAlfaComponent() component.ApiLambdaDynamoDb
	EchoAppAlfaStage() awscdk.Stage
}

func NewEchoAppAlfa(scope constructs.Construct, id *string, props *EchoAppAlfaProps) EchoAppAlfa {

	var sprops EchoAppAlfaProps = EchoAppAlfaProps_DEFAULT
	var sid EchoAppAlfaIds = sprops.EchoAppAlfaIds

	if props != nil {
		sprops = *props
		sid = sprops.EchoAppAlfaIds
	}

	if id != nil {
		sid.EchoAppAlfa_Id = *id
	}

	stage := awscdk.NewStage(scope, jsii.String(sid.EchoAppAlfa_Id), &sprops.StageProps)

	cpt := component.NewApiLambdaDynamoDb(stage, jsii.String(sid.ApiLambdaDynamoDb_Id),
		&sprops.ApiLambdaDynamoDbProps)

	return echoAppAlfa{stage, cpt}
}

// CONFIGURATIONS
var EchoAppAlfaProps_DEFAULT EchoAppAlfaProps = EchoAppAlfaProps{
	StageProps:             environment.Stage_DEFAULT,
	ApiLambdaDynamoDbProps: component.ApiLambdaDynamoDbProps_DEFAULT,
	EchoAppAlfaIds: EchoAppAlfaIds{
		EchoAppAlfa_Id:       "EchoApp-ALFA-Implementation-default",
		ApiLambdaDynamoDb_Id: "EchoApp-Component-ALFA-default",
	},
}

var EchoAppAlfaProps_DEV EchoAppAlfaProps = EchoAppAlfaProps{
	StageProps:             environment.Stage_DEV,
	ApiLambdaDynamoDbProps: component.ApiLambdaDynamoDbProps_DEV,
	EchoAppAlfaIds: EchoAppAlfaIds{
		EchoAppAlfa_Id:       "EchoApp-ALFA-Implementation-dev",
		ApiLambdaDynamoDb_Id: "EchoApp-Component-ALFA-dev",
	},
}

var EchoAppAlfaProps_PROD EchoAppAlfaProps = EchoAppAlfaProps{
	StageProps:             environment.Stage_PROD,
	ApiLambdaDynamoDbProps: component.ApiLambdaDynamoDbProps_PROD,
	EchoAppAlfaIds: EchoAppAlfaIds{
		EchoAppAlfa_Id:       "EchoApp-ALFA-Implementation-prod",
		ApiLambdaDynamoDb_Id: "EchoApp-Component-ALFA-prod",
	},
}
