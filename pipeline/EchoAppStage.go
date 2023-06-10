package pipeline

import (
	"writer_storage_app/component"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type EchoAppPipelineStageProps struct {
	stageprops *awscdk.StageProps
	CptProps   *awscdk.StackProps
	Cpt        component.Component
}

func EchoAppPipelineStage(scope constructs.Construct, id *string, props *EchoAppPipelineStageProps) awscdk.Stage {
	var sprops awscdk.StageProps
	if props != nil {
		sprops = *props.stageprops
	}
	stage := awscdk.NewStage(scope, id, &sprops)
	//uncouple component
	props.Cpt.NewComponentStack(stage, id,
		*props.CptProps)

	return stage
}
