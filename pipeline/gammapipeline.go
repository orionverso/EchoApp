package pipeline

import (
	"writer_storage_app/environment"
	"writer_storage_app/pipeline/stages"
<<<<<<< Updated upstream
=======
	"writer_storage_app/pipeline/stages/auxiliar"
>>>>>>> Stashed changes

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarconnections"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type GammaPipelineIds struct {
	GammaPipeline_Id                        string
	CfnConnection_Id                        string
	CodePipelineSource_Connection_Id        string
	CodePipelineSource_Connection_branch_Id string
	CodeBuildStep_Id                        string
	CodePipeline_Id                         string
	EchoAppGamma_Id                         string
}

type GammaPipelineProps struct {
	StackProps              awscdk.StackProps
	CfnConnectionProps      awscodestarconnections.CfnConnectionProps
	ConnectionSourceOptions pipelines.ConnectionSourceOptions
	CodeBuildStepProps      pipelines.CodeBuildStepProps
	CodePipelineProps       pipelines.CodePipelineProps
<<<<<<< Updated upstream
=======
	EcrRepositoryProps      auxiliar.EcrRepositoryProps
>>>>>>> Stashed changes
	EchoAppGammaProps_1ENV  stages.EchoAppGammaProps
	EchoAppGammaProps_2ENV  stages.EchoAppGammaProps
	AddStageOpts            pipelines.AddStageOpts
	//Identifiers
	GammaPipelineIds
}

type gammaPipeline struct {
	awscdk.Stack
	cfnConnection      awscodestarconnections.CfnConnection
	codePipelineSource pipelines.CodePipelineSource
	codeBuildStep      pipelines.CodeBuildStep
	codePipeline       pipelines.CodePipeline
	echoAppGamma_1ENV  stages.EchoAppGamma
	echoAppGamma_2ENV  stages.EchoAppGamma
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

func (af gammaPipeline) GammaEchoAppGamma_1ENV() stages.EchoAppGamma {
	return af.echoAppGamma_1ENV
}

func (af gammaPipeline) GammaEchoAppGamma_2ENV() stages.EchoAppGamma {
	return af.echoAppGamma_2ENV
}

type GammaPipeline interface {
	awscdk.Stack
	GammaCfnConnection() awscodestarconnections.CfnConnection
	GammaCodePipelineSource() pipelines.CodePipelineSource
	GammaCodeBuildStep() pipelines.CodeBuildStep
	GammaCodePipeline() pipelines.CodePipeline
	GammaEchoAppGamma_1ENV() stages.EchoAppGamma
	GammaEchoAppGamma_2ENV() stages.EchoAppGamma
}

func NewGammaPipeline(scope constructs.Construct, id *string, props *GammaPipelineProps) GammaPipeline {

	var sprops GammaPipelineProps = GammaPipeline_DEFAULT
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

	deploy_First_Env := stages.NewEchoAppGamma(stack, nil, &sprops.EchoAppGammaProps_1ENV) // Development Environment
	pipe.AddStage(deploy_First_Env.EchoAppGammaStage(), nil)

	deploy_Second_Env := stages.NewEchoAppGamma(stack, nil, &sprops.EchoAppGammaProps_2ENV) //Staging Environment
	pipe.AddStage(deploy_Second_Env.EchoAppGammaStage(), &sprops.AddStageOpts)

	return gammaPipeline{stack, conn, GithubRepository, Template, pipe, deploy_First_Env, deploy_Second_Env}
}

// CONFIGURATIONS
var GammaPipeline_DEFAULT GammaPipelineProps = GammaPipelineProps{

	StackProps: environment.StackProps_DEFAULT,

	CfnConnectionProps: awscodestarconnections.CfnConnectionProps{
		ConnectionName: jsii.String("GithubConnection"),
		ProviderType:   jsii.String("GitHub"),
	},

	CodeBuildStepProps: pipelines.CodeBuildStepProps{
		Commands: jsii.Strings("npm install -g aws-cdk", "cd webserver", "docker build -t writer-server .", "cd ..", "cdk synth"),
		Env: &map[string]*string{
			"CDK_DEFAULT_REGION":  environment.StackProps_DEFAULT.Env.Region,
			"CDK_DEFAULT_ACCOUNT": environment.StackProps_DEFAULT.Env.Account,
			"CDK_DEV_REGION":      environment.StackProps_DEV.Env.Region,
			"CDK_DEV_ACCOUNT":     environment.StackProps_DEV.Env.Account,
		},
	},

	CodePipelineProps: pipelines.CodePipelineProps{
		PipelineName:     jsii.String("EchoAppGamma-Pipeline-default"),
		CrossAccountKeys: jsii.Bool(true),
	},

<<<<<<< Updated upstream
=======
	EcrRepositoryProps: auxiliar.EcrRepositoryProps_DEFAULT,

>>>>>>> Stashed changes
	EchoAppGammaProps_1ENV: stages.EchoAppGammaProps_DEFAULT,
	EchoAppGammaProps_2ENV: stages.EchoAppGammaProps_DEV,

	AddStageOpts: pipelines.AddStageOpts{
		Pre: &[]pipelines.Step{
			pipelines.NewManualApprovalStep(jsii.String("PromoteComponentToProduction"), &pipelines.ManualApprovalStepProps{
				Comment: jsii.String("LAST CHECK BEFORE PRODUCTION"),
			}),
		},
	},

	GammaPipelineIds: GammaPipelineIds{
		GammaPipeline_Id: "GammaPipeline-default",
		CfnConnection_Id: "CodestarConnectionToGithub",

		CodePipelineSource_Connection_Id:        "orionverso/EchoApp_mock",
		CodePipelineSource_Connection_branch_Id: "gamma",
		CodeBuildStep_Id:                        "SynthStep",
		CodePipeline_Id:                         "EchoAppGamma-Pipeline",
		EchoAppGamma_Id:                         "DeployStageOf-EchoAppGamma",
	},
}

var GammaPipeline_DEV GammaPipelineProps = GammaPipelineProps{

	StackProps: environment.StackProps_DEV,

	CfnConnectionProps: awscodestarconnections.CfnConnectionProps{
		ConnectionName: jsii.String("GithubConnection"),
		ProviderType:   jsii.String("GitHub"),
	},

	CodeBuildStepProps: pipelines.CodeBuildStepProps{
		Commands: jsii.Strings("npm install -g aws-cdk", "cd webserver", "docker build -t writer-server .", "cd ..", "cdk synth"),
		Env: &map[string]*string{
			"CDK_DEV_REGION":   environment.StackProps_DEV.Env.Region,
			"CDK_DEV_ACCOUNT":  environment.StackProps_DEV.Env.Account,
			"CDK_PROD_REGION":  environment.StackProps_PROD.Env.Region,
			"CDK_PROD_ACCOUNT": environment.StackProps_PROD.Env.Account,
		},
	},

	CodePipelineProps: pipelines.CodePipelineProps{
		PipelineName:     jsii.String("EchoAppGamma-Pipeline-dev"),
		CrossAccountKeys: jsii.Bool(true),
	},

<<<<<<< Updated upstream
=======
	EcrRepositoryProps: auxiliar.EcrRepositoryProps_DEV,

>>>>>>> Stashed changes
	EchoAppGammaProps_1ENV: stages.EchoAppGammaProps_DEV,
	EchoAppGammaProps_2ENV: stages.EchoAppGammaProps_PROD,

	AddStageOpts: pipelines.AddStageOpts{
		Pre: &[]pipelines.Step{
			pipelines.NewManualApprovalStep(jsii.String("PromoteComponentToProduction"), &pipelines.ManualApprovalStepProps{
				Comment: jsii.String("LAST CHECK BEFORE PRODUCTION"),
			}),
		},
	},

	GammaPipelineIds: GammaPipelineIds{
		GammaPipeline_Id:                        "GammaPipeline-dev",
		CfnConnection_Id:                        "CodestarConnectionToGithub",
		CodePipelineSource_Connection_Id:        "orionverso/EchoApp_mock",
		CodePipelineSource_Connection_branch_Id: "gamma",
		CodeBuildStep_Id:                        "SynthStep",
		CodePipeline_Id:                         "EchoAppGamma-Pipeline",
		EchoAppGamma_Id:                         "DeployStageOf-EchoAppGamma",
	},
}
