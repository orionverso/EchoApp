package pipeline

import (
	"writer_storage_app/pipeline/stages"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarconnections"
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

	sprops.CodeBuildStepProps.Input = GithubRepository

	Template := pipelines.NewCodeBuildStep(
		jsii.String(sid.CodeBuildStep_Id),
		&sprops.CodeBuildStepProps,
	)

	sprops.CodePipelineProps.Synth = Template

	pipe := pipelines.NewCodePipeline(stack, jsii.String(sid.CodePipeline_Id), &sprops.CodePipelineProps)
	//First Environment: Dev
	deploy_FIRST_ENV := stages.NewEchoAppGamma(stack, nil, &sprops.EchoAppGammaProps_FIRST_ENV) // Development Environment

	sprops.AddedStep.CheckPushImageStep.AddStepDependency(sprops.AddedStep.PushImageStep)

	RepoStackPosition, FargateStackPosition := 0, 1
	addStackToStackSteps(deploy_FIRST_ENV.EchoAppGammaRepositoryComponentStack(), RepoStackPosition, &sprops.AddStageOpts_FIRST_ENV)
	addStackToStackSteps(deploy_FIRST_ENV.EchoAppGammaFargateS3ComponentStack(), FargateStackPosition, &sprops.AddStageOpts_FIRST_ENV)

	pipe.AddStage(deploy_FIRST_ENV.EchoAppGammaStage(), &sprops.AddStageOpts_FIRST_ENV)

	//Prepare second Enviroment

	deploy_preparation := stages.NewNextDeployPreparation(stack, nil, &sprops.NextDeployPreparationProps)

	addStackToStackSteps(deploy_preparation.RolePushImageCrossAccount().RolePushImageCrossAccountStack(), 0, &sprops.AddStageOpts_NEXT_ENV_PREP)

	pipe.AddStage(deploy_preparation.NextDeployPreparationStage(), &sprops.AddStageOpts_NEXT_ENV_PREP)

	//Second Enviroment: Prod
	deploy_SECOND_ENV := stages.NewEchoAppGamma(stack, nil, &sprops.EchoAppGammaProps_SECOND_ENV) //PROD Environment

	sprops.AddedStep = AddedStep_PROD //Change to Prods steps

	sprops.AddedStep.CheckPushImageStep.AddStepDependency(sprops.AddedStep.PushImageStep)

	addStackToStackSteps(deploy_SECOND_ENV.EchoAppGammaRepositoryComponentStack(), RepoStackPosition, &sprops.AddStageOpts_SECOND_ENV) // 0

	pipe.AddStage(deploy_SECOND_ENV.EchoAppGammaStage(), &sprops.AddStageOpts_SECOND_ENV)

	return gammaPipeline{stack, conn, GithubRepository, Template, pipe, deploy_FIRST_ENV, deploy_SECOND_ENV}
}
