package pipeline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/jsii-runtime-go"
)

//Helper is for passing parameters to configurations at runtime that cannot be done in one step.

func addStackToStackSteps(stack awscdk.Stack, pos int, props *pipelines.AddStageOpts) {
	//Remember slice pointer behavior
	var StackStepsPtr []*pipelines.StackSteps = *props.StackSteps
	var StackStepPtr *pipelines.StackSteps = StackStepsPtr[pos]
	StackStepPtr.Stack = stack
}

func pushImagePolicy_DEV() awsiam.PolicyStatement {

	pl := awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Effect:    awsiam.Effect_ALLOW,
		Resources: jsii.Strings("*"),
		Actions: jsii.Strings("ecr:GetAuthorizationToken",
			"ecr:BatchCheckLayerAvailability",
			"ecr:GetDownloadUrlForLayer",
			"ecr:GetRepositoryPolicy",
			"ecr:DescribeRepositories",
			"ecr:ListImages",
			"ecr:DescribeImages",
			"ecr:BatchGetImage",
			"ecr:InitiateLayerUpload",
			"ecr:UploadLayerPart",
			"ecr:CompleteLayerUpload",
			"ecr:PutImage",
		),
	})

	return pl
}
