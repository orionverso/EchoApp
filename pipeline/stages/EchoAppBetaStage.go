package stages

import (
	"writer_storage_app/component"
	"writer_storage_app/environment"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type EchoAppBetaIds struct {
	EchoAppBeta_Id string
	ApiLambdaS3_Id string
}

type EchoAppBetaProps struct {
	StageProps       awscdk.StageProps
	ApiLambdaS3Props component.ApiLambdaS3Props
	//Identifiers
	EchoAppBetaIds
}

type echoAppBeta struct {
	awscdk.Stage
	cpt component.ApiLambdaS3
}

func (ec echoAppBeta) EchoAppBetaComponent() component.ApiLambdaS3 {
	return ec.cpt
}

func (ec echoAppBeta) EchoAppBetaStage() awscdk.Stage {
	return ec.Stage
}

type EchoAppBeta interface {
	awscdk.Stage
	EchoAppBetaComponent() component.ApiLambdaS3
	EchoAppBetaStage() awscdk.Stage
}

func NewEchoAppBeta(scope constructs.Construct, id *string, props *EchoAppBetaProps) EchoAppBeta {

	var sprops EchoAppBetaProps = EchoAppBetaProps_DEFAULT
	var sid EchoAppBetaIds = sprops.EchoAppBetaIds

	if props != nil {
		sprops = *props
		sid = sprops.EchoAppBetaIds
	}

	if id != nil {
		sid.EchoAppBeta_Id = *id
	}

	stage := awscdk.NewStage(scope, jsii.String(sid.ApiLambdaS3_Id), &sprops.StageProps)

	cpt := component.NewApiLambdaS3(stage, jsii.String(sid.ApiLambdaS3_Id),
		&sprops.ApiLambdaS3Props)

	return echoAppBeta{stage, cpt}
}

// CONFIGURATIONS
var EchoAppBetaProps_DEFAULT EchoAppBetaProps = EchoAppBetaProps{
	StageProps:       environment.Stage_DEFAULT,
	ApiLambdaS3Props: component.ApiLambdaS3Props_DEFAULT,
	EchoAppBetaIds: EchoAppBetaIds{
		EchoAppBeta_Id: "EchoApp-BETA-Implementation-default",
		ApiLambdaS3_Id: "EchoApp-Component-BETA-default",
	},
}

var EchoAppBetaProps_DEV EchoAppBetaProps = EchoAppBetaProps{
	StageProps:       environment.Stage_DEV,
	ApiLambdaS3Props: component.ApiLambdaS3Props_DEV,
	EchoAppBetaIds: EchoAppBetaIds{
		EchoAppBeta_Id: "EchoApp-BETA-Implementation-dev",
		ApiLambdaS3_Id: "EchoApp-Component-BETA-dev",
	},
}

var EchoAppBetaProps_PROD EchoAppBetaProps = EchoAppBetaProps{
	StageProps:       environment.Stage_PROD,
	ApiLambdaS3Props: component.ApiLambdaS3Props_PROD,
	EchoAppBetaIds: EchoAppBetaIds{
		EchoAppBeta_Id: "EchoApp-BETA-Implementation-prod",
		ApiLambdaS3_Id: "EchoApp-Component-BETA-prod",
	},
}
