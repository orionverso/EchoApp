package stages

import (
	"writer_storage_app/component"
	"writer_storage_app/environment"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type EchoAppGammaIds struct {
	EchoAppGamma_Id string
	FargateS3_Id    string
}

type EchoAppGammaProps struct {
	StageProps     awscdk.StageProps
	FargateS3Props component.FargateS3Props
	//Identifiers
	EchoAppGammaIds
}

type echoAppGamma struct {
	awscdk.Stage
	cpt component.FargateS3
}

func (ec echoAppGamma) EchoAppGammaComponent() component.FargateS3 {
	return ec.cpt
}

func (ec echoAppGamma) EchoAppGammaStage() awscdk.Stage {
	return ec.Stage
}

type EchoAppGamma interface {
	awscdk.Stage
	EchoAppGammaComponent() component.FargateS3
	EchoAppGammaStage() awscdk.Stage
}

func NewEchoAppGamma(scope constructs.Construct, id *string, props *EchoAppGammaProps) EchoAppGamma {

	var sprops EchoAppGammaProps = EchoAppGammaProps_DEFAULT
	var sid EchoAppGammaIds = sprops.EchoAppGammaIds

	if props != nil {
		sprops = *props
		sid = sprops.EchoAppGammaIds
	}

	if id != nil {
		sid.EchoAppGamma_Id = *id
	}

	stage := awscdk.NewStage(scope, jsii.String(sid.EchoAppGamma_Id), &sprops.StageProps)

	cpt := component.NewFargateS3(stage, jsii.String(sid.FargateS3_Id),
		&sprops.FargateS3Props)

	return echoAppGamma{stage, cpt}
}

// CONFIGURATIONS
var EchoAppGammaProps_DEFAULT EchoAppGammaProps = EchoAppGammaProps{
	StageProps:     environment.Stage_DEFAULT,
	FargateS3Props: component.FargateS3Props_DEFAULT,
	EchoAppGammaIds: EchoAppGammaIds{
		EchoAppGamma_Id: "EchoApp-GAMMA-Implementation-default",
		FargateS3_Id:    "EchoApp-Component-GAMMA-default",
	},
}

var EchoAppGammaProps_DEV EchoAppGammaProps = EchoAppGammaProps{
	StageProps:     environment.Stage_DEV,
	FargateS3Props: component.FargateS3Props_DEV,
	EchoAppGammaIds: EchoAppGammaIds{
		EchoAppGamma_Id: "EchoApp-GAMMA-Implementation-dev",
		FargateS3_Id:    "EchoApp-Component-GAMMA-dev",
	},
}

var EchoAppGammaProps_PROD EchoAppGammaProps = EchoAppGammaProps{
	StageProps:     environment.Stage_PROD,
	FargateS3Props: component.FargateS3Props_PROD,
	EchoAppGammaIds: EchoAppGammaIds{
		EchoAppGamma_Id: "EchoApp-GAMMA-Implementation-prod",
		FargateS3_Id:    "EchoApp-Component-GAMMA-prod",
	},
}
