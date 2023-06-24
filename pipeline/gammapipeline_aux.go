package pipeline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
)

//Auxiliar is for passing parameters to configurations at runtime that cannot be done in one step.

// Add stack at runtime to CodePipelineProps. The idea is not to overwrite the settings
func addStackToStackSteps(stack awscdk.Stack, pos int, props *pipelines.AddStageOpts) {
	//Remember slice pointer behavior
	var StackStepsPtr []*pipelines.StackSteps = *props.StackSteps
	var StackStepPtr *pipelines.StackSteps = StackStepsPtr[pos]
	StackStepPtr.Stack = stack
}
