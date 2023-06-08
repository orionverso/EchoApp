package pipeline

import (
	"writer_storage_app/component"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type EchoAppPipelineStageProps struct {
	awscdk.StageProps
	Cpt component.Component
}

func EchoAppPipelineStage(scope constructs.Construct, id *string, props *EchoAppPipelineStageProps) awscdk.Stage {
	var sprops awscdk.StageProps
	if props != nil {
		sprops = props.StageProps
	}
	stage := awscdk.NewStage(scope, id, &sprops)
	//uncouple
	props.Cpt.NewComponentStack(stage, id, awscdk.StackProps{})

	return stage
}
