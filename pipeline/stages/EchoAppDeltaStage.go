package stages

import (
	"writer_storage_app/component"
	"writer_storage_app/environment"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type EchoAppDeltaIds struct {
	EchoAppDelta_Id    string
	FargateDynamoDb_Id string
}

type EchoAppDeltaProps struct {
	StageProps           awscdk.StageProps
	FargateDynamoDbProps component.FargateDynamoDbProps
	//Identifiers
	EchoAppDeltaIds
}

type echoAppDelta struct {
	awscdk.Stage
	cpt component.FargateDynamoDb
}

func (ec echoAppDelta) EchoAppDeltaComponent() component.FargateDynamoDb {
	return ec.cpt
}

func (ec echoAppDelta) EchoAppDeltaStage() awscdk.Stage {
	return ec.Stage
}

type EchoAppDelta interface {
	awscdk.Stage
	EchoAppDeltaComponent() component.FargateDynamoDb
	EchoAppDeltaStage() awscdk.Stage
}

func NewEchoAppDelta(scope constructs.Construct, id *string, props *EchoAppDeltaProps) EchoAppDelta {

	var sprops EchoAppDeltaProps = EchoAppDeltaProps_DEFAULT
	var sid EchoAppDeltaIds = sprops.EchoAppDeltaIds

	if props != nil {
		sprops = *props
		sid = sprops.EchoAppDeltaIds
	}

	if id != nil {
		sid.EchoAppDelta_Id = *id
	}

	stage := awscdk.NewStage(scope, jsii.String(sid.EchoAppDelta_Id), &sprops.StageProps)

	cpt := component.NewFargateDynamoDb(stage, jsii.String(sid.FargateDynamoDb_Id),
		&sprops.FargateDynamoDbProps)

	return echoAppDelta{stage, cpt}
}

// CONFIGURATIONS
var EchoAppDeltaProps_DEFAULT EchoAppDeltaProps = EchoAppDeltaProps{
	StageProps:           environment.Stage_DEFAULT,
	FargateDynamoDbProps: component.FargateDynamoDbProps_DEFAULT,
	EchoAppDeltaIds: EchoAppDeltaIds{
		EchoAppDelta_Id:    "EchoAppDelta-DELTA-Implementaion-default",
		FargateDynamoDb_Id: "EchoAppDelta-Component-DELTA-default",
	},
}

var EchoAppDeltaProps_DEV EchoAppDeltaProps = EchoAppDeltaProps{
	StageProps:           environment.Stage_DEV,
	FargateDynamoDbProps: component.FargateDynamoDbProps_DEV,
	EchoAppDeltaIds: EchoAppDeltaIds{
		EchoAppDelta_Id:    "EchoAppDelta-DELTA-Implementaion-dev",
		FargateDynamoDb_Id: "EchoAppDelta-Component-DELTA-dev",
	},
}

var EchoAppDeltaProps_PROD EchoAppDeltaProps = EchoAppDeltaProps{
	StageProps:           environment.Stage_PROD,
	FargateDynamoDbProps: component.FargateDynamoDbProps_PROD,
	EchoAppDeltaIds: EchoAppDeltaIds{
		EchoAppDelta_Id:    "EchoAppDelta-DELTA-Implementaion-prod",
		FargateDynamoDb_Id: "EchoAppDelta-Component-DELTA-prod",
	},
}
