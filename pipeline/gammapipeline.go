package pipeline

import (
	"writer_storage_app/pipeline/stages"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarconnections"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type GammaPipeline interface {
	awscdk.Stack
	GammaCfnConnection() awscodestarconnections.CfnConnection
	GammaCodePipelineSource() pipelines.CodePipelineSource
	GammaCodeBuildStep() pipelines.CodeBuildStep
	GammaCodePipeline() pipelines.CodePipeline
	GammaEchoAppGamma_FIRST_ENV() stages.EchoAppGamma
	GammaEchoAppGamma_SECOND_ENV() stages.EchoAppGamma
}

type gammaPipeline struct {
	awscdk.Stack
	cfnConnection           awscodestarconnections.CfnConnection
	codePipelineSource      pipelines.CodePipelineSource
	codeBuildStep           pipelines.CodeBuildStep
	codePipeline            pipelines.CodePipeline
	echoAppGamma_FIRST_ENV  stages.EchoAppGamma
	echoAppGamma_SECOND_ENV stages.EchoAppGamma
}

func (af gammaPipeline) GammaCfnConnection() awscodestarconnections.CfnConnection {
	return af.cfnConnection
}

func (af gammaPipeline) GammaCodePipelineSource() pipelines.CodePipelineSource {
	return af.codePipelineSource
}

func (af gammaPipeline) GammaCodeBuildStep() pipelines.CodeBuildStep {
	return af.codeBuildStep
}

func (af gammaPipeline) GammaCodePipeline() pipelines.CodePipeline {
	return af.codePipeline
}

func (af gammaPipeline) GammaEchoAppGamma_FIRST_ENV() stages.EchoAppGamma {
	return af.echoAppGamma_FIRST_ENV
}

func (af gammaPipeline) GammaEchoAppGamma_SECOND_ENV() stages.EchoAppGamma {
	return af.echoAppGamma_SECOND_ENV
}

func NewGammaPipeline(scope constructs.Construct, id *string, props *GammaPipelineProps) GammaPipeline {

	var sprops GammaPipelineProps = GammaPipelineProps_DEV
	var sid GammaPipelineIds = sprops.GammaPipelineIds

	if props != nil {
		sprops = *props
		sid = sprops.GammaPipelineIds
	}

	if id != nil {
		sid.GammaPipeline_Id = *id
	}

	stack := awscdk.NewStack(scope, jsii.String(sid.GammaPipeline_Id), &sprops.StackProps)
	//Connect to GitHub
	//You must accepted connection manually
	//https://docs.aws.amazon.com/codepipeline/latest/userguide/connections-github.html#connections-github-cli
	conn := awscodestarconnections.NewCfnConnection(stack, jsii.String(sid.CfnConnection_Id), &sprops.CfnConnectionProps)

	sprops.ConnectionSourceOptions.ConnectionArn = conn.AttrConnectionArn()

	GithubRepository := pipelines.CodePipelineSource_Connection(
		jsii.String(sid.CodePipelineSource_Connection_Id),
		jsii.String(sid.CodePipelineSource_Connection_branch_Id),
		&sprops.ConnectionSourceOptions,
	)

	sprops.CodeBuildSynthStepProps.Input = GithubRepository

	Template := pipelines.NewCodeBuildStep(
		jsii.String(sid.CodeBuildStep_Id),
		&sprops.CodeBuildSynthStepProps,
	)

	sprops.CodePipelineProps.Synth = Template

	pipe := pipelines.NewCodePipeline(stack, jsii.String(sid.CodePipeline_Id), &sprops.CodePipelineProps)

	//First Environment: Dev

	deploy_FIRST_ENV := stages.NewEchoAppGamma(stack, nil, &sprops.EchoAppGammaProps_FIRST_ENV) // Development Environment

	pushImageToRepo_v1 := pipelines.NewCodeBuildStep(jsii.String(sid.PushImageStep_Id), &sprops.AddedStepProps.PushImageStepProps)

	checkImageWasPushed_v1 := pipelines.NewManualApprovalStep(jsii.String(sid.CheckPushImageStep_Id), &sprops.AddedStepProps.CheckPushImageStepProps)

	checkImageWasPushed_v1.AddStepDependency(pushImageToRepo_v1)

	repoSteps_v1 := pipelines.StackSteps{
		Stack: deploy_FIRST_ENV.EchoAppGammaRepositoryComponentStack(),
		Post:  &[]pipelines.Step{pushImageToRepo_v1, checkImageWasPushed_v1},
	}

	*sprops.AddStageOpts_FIRST_ENV.StackSteps = append(*sprops.AddStageOpts_FIRST_ENV.StackSteps, &repoSteps_v1)

	//pipe.AddStage(deploy_FIRST_ENV.EchoAppGammaStage(), &sprops.AddStageOpts_FIRST_ENV) Now, we dont need test this part

	//Preparation Stage

	deploy_preparation := stages.NewNextDeployPreparation(stack, nil, &sprops.NextDeployPreparationProps)

	sprops.AddedStepProps = AddedStepProps_PROD
	sid.AddedStepIds = AddedStepProps_PROD.AddedStepIds

	promoteProdDecision := pipelines.NewManualApprovalStep(jsii.String(sid.PromoteToProduction_Id), &sprops.AddedStepProps.PromoteToProductionProps)

	//It is more easy with stackdeployment and stagedeployment. Explore it!

	RoleSteps := pipelines.StackSteps{ //It is more easy with stackdeployment and stagedeployment
		Stack: deploy_preparation.RolePushImageCrossAccount().RolePushImageCrossAccountStack(),
		Post:  &[]pipelines.Step{promoteProdDecision},
	}

	*sprops.AddStageOpts_NEXT_ENV_PREP.StackSteps = append(*sprops.AddStageOpts_NEXT_ENV_PREP.StackSteps, &RoleSteps)

	pipe.AddStage(deploy_preparation.NextDeployPreparationStage(), &sprops.AddStageOpts_NEXT_ENV_PREP)

	//Second Environment: Prod

	deploy_SECOND_ENV := stages.NewEchoAppGamma(stack, nil, &sprops.EchoAppGammaProps_SECOND_ENV) //PROD Environment

	*sprops.AddedStepProps.PushImageStepProps.RolePolicyStatements = append(*sprops.AddedStepProps.PushImageStepProps.RolePolicyStatements,
		awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
			Effect:    awsiam.Effect_ALLOW,
			Actions:   jsii.Strings("sts:AssumeRole"),
			Resources: jsii.Strings(*deploy_preparation.RolePushImageCrossAccount().RolePushImageCrossAccountRole().RoleArn()),
		}),
	)

	sprops.AddedStepProps.AddEnvVar(jsii.String("PUSH_ROLE_ARN"), deploy_preparation.RolePushImageCrossAccount().RolePushImageCrossAccountRole().RoleArn(), &sprops.AddedStepProps.PushImageStepProps)

	pushImageToRepo_v2 := pipelines.NewCodeBuildStep(jsii.String(sid.PushImageStep_Id), &sprops.AddedStepProps.PushImageStepProps)

	checkImageWasPushed_v2 := pipelines.NewManualApprovalStep(jsii.String(sid.CheckPushImageStep_Id), &sprops.AddedStepProps.CheckPushImageStepProps)

	checkImageWasPushed_v2.AddStepDependency(pushImageToRepo_v2)

	repoSteps_v2 := pipelines.StackSteps{
		Stack: deploy_SECOND_ENV.EchoAppGammaRepositoryComponentStack(),
		Post:  &[]pipelines.Step{pushImageToRepo_v2, checkImageWasPushed_v2},
	}

	*sprops.AddStageOpts_SECOND_ENV.StackSteps = append(*sprops.AddStageOpts_SECOND_ENV.StackSteps, &repoSteps_v2)

	pipe.AddStage(deploy_SECOND_ENV.EchoAppGammaStage(), &sprops.AddStageOpts_SECOND_ENV)

	return gammaPipeline{stack, conn, GithubRepository, Template, pipe, deploy_FIRST_ENV, deploy_SECOND_ENV}
}
