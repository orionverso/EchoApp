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
	GammaEchoAppGamma_First_Env() stages.EchoAppGamma
	GammaEchoAppGamma_Second_Env() stages.EchoAppGamma
}

type gammaPipeline struct {
	awscdk.Stack
	cfnConnection           awscodestarconnections.CfnConnection
	codePipelineSource      pipelines.CodePipelineSource
	codeBuildStep           pipelines.CodeBuildStep
	codePipeline            pipelines.CodePipeline
	echoAppGamma_First_Env  stages.EchoAppGamma
	echoAppGamma_Second_Env stages.EchoAppGamma
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

func (af gammaPipeline) GammaEchoAppGamma_First_Env() stages.EchoAppGamma {
	return af.echoAppGamma_First_Env
}

func (af gammaPipeline) GammaEchoAppGamma_Second_Env() stages.EchoAppGamma {
	return af.echoAppGamma_Second_Env
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
	deploy_First_Env := stages.NewEchoAppGamma(stack, nil, &sprops.EchoAppGammaProps_First_Env) // Development Environment

	sprops.AddedStep.CheckPushImageStep.AddStepDependency(sprops.AddedStep.PushImageStep)

	RepoStackPosition, FargateStackPosition := 0, 1
	addStackToStackSteps(deploy_First_Env.EchoAppGammaRepositoryComponentStack(), RepoStackPosition, &sprops.AddStageOpts_First_Env)
	addStackToStackSteps(deploy_First_Env.EchoAppGammaFargateS3ComponentStack(), FargateStackPosition, &sprops.AddStageOpts_First_Env)

	pipe.AddStage(deploy_First_Env.EchoAppGammaStage(), &sprops.AddStageOpts_First_Env)
	//Second Enviroment: Prod
	deploy_Second_Env := stages.NewEchoAppGamma(stack, nil, &sprops.EchoAppGammaProps_Second_Env) //PROD Environment

	sprops.AddedStep = AddedStep_PROD //Change to Prods steps

	sprops.AddedStep.CheckPushImageStep.AddStepDependency(sprops.AddedStep.PushImageStep)

	addStackToStackSteps(deploy_Second_Env.EchoAppGammaRepositoryComponentStack(), RepoStackPosition, &sprops.AddStageOpts_Second_Env) // 0

	pipe.AddStage(deploy_Second_Env.EchoAppGammaStage(), &sprops.AddStageOpts_Second_Env)

	return gammaPipeline{stack, conn, GithubRepository, Template, pipe, deploy_First_Env, deploy_Second_Env}
}
