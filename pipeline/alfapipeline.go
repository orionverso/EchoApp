package pipeline

import (
	"writer_storage_app/environment"
	"writer_storage_app/pipeline/stages"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarconnections"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type AlfaPipelineIds struct {
	AlfaPipeline_Id                         string
	CfnConnection_Id                        string
	CodePipelineSource_Connection_Id        string
	CodePipelineSource_Connection_branch_Id string
	CodeBuildStep_Id                        string
	CodePipeline_Id                         string
	EchoAppAlfa_Id                          string
}

type AlfaPipelineProps struct {
	StackProps              awscdk.StackProps
	CfnConnectionProps      awscodestarconnections.CfnConnectionProps
	ConnectionSourceOptions pipelines.ConnectionSourceOptions
	CodeBuildStepProps      pipelines.CodeBuildStepProps
	CodePipelineProps       pipelines.CodePipelineProps
	EchoAppAlfaProps_1ENV   stages.EchoAppAlfaProps
	EchoAppAlfaProps_2ENV   stages.EchoAppAlfaProps
	EchoAppAlfaProps_3ENV   stages.EchoAppAlfaProps
	AddStageOpts            pipelines.AddStageOpts
	//Identifiers
	AlfaPipelineIds
}

type alfaPipeline struct {
	awscdk.Stack
	cfnConnection      awscodestarconnections.CfnConnection
	codePipelineSource pipelines.CodePipelineSource
	codeBuildStep      pipelines.CodeBuildStep
	codePipeline       pipelines.CodePipeline
	echoAppAlfa_1ENV   stages.EchoAppAlfa
	echoAppAlfa_2ENV   stages.EchoAppAlfa
}

func (af alfaPipeline) AlfaCfnConnection() awscodestarconnections.CfnConnection {
	return af.cfnConnection
}

func (af alfaPipeline) AlfaCodePipelineSource() pipelines.CodePipelineSource {
	return af.codePipelineSource
}

func (af alfaPipeline) AlfaCodeBuildStep() pipelines.CodeBuildStep {
	return af.codeBuildStep
}

func (af alfaPipeline) AlfaCodePipeline() pipelines.CodePipeline {
	return af.codePipeline
}

func (af alfaPipeline) AlfaEchoAppAlfa_1ENV() stages.EchoAppAlfa {
	return af.echoAppAlfa_1ENV
}

func (af alfaPipeline) AlfaEchoAppAlfa_2ENV() stages.EchoAppAlfa {
	return af.echoAppAlfa_2ENV
}

type AlfaPipeline interface {
	awscdk.Stack
	AlfaCfnConnection() awscodestarconnections.CfnConnection
	AlfaCodePipelineSource() pipelines.CodePipelineSource
	AlfaCodeBuildStep() pipelines.CodeBuildStep
	AlfaCodePipeline() pipelines.CodePipeline
	AlfaEchoAppAlfa_1ENV() stages.EchoAppAlfa
	AlfaEchoAppAlfa_2ENV() stages.EchoAppAlfa
}

func NewAlfaPipeline(scope constructs.Construct, id *string, props *AlfaPipelineProps) AlfaPipeline {

	var sprops AlfaPipelineProps = AlfaPipeline_DEFAULT
	var sid AlfaPipelineIds = sprops.AlfaPipelineIds

	if props != nil {
		sprops = *props
		sid = sprops.AlfaPipelineIds
	}

	if id != nil {
		sid.AlfaPipeline_Id = *id
	}

	stack := awscdk.NewStack(scope, jsii.String(sid.AlfaPipeline_Id), &sprops.StackProps)
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

	deploy_First_Env := stages.NewEchoAppAlfa(stack, nil, &sprops.EchoAppAlfaProps_1ENV) // Development Environment
	pipe.AddStage(deploy_First_Env.EchoAppAlfaStage(), nil)

	deploy_Second_Env := stages.NewEchoAppAlfa(stack, nil, &sprops.EchoAppAlfaProps_2ENV) //Staging Environment
	pipe.AddStage(deploy_Second_Env.EchoAppAlfaStage(), &sprops.AddStageOpts)

	return alfaPipeline{stack, conn, GithubRepository, Template, pipe, deploy_First_Env, deploy_Second_Env}
}

// CONFIGURATIONS
var AlfaPipeline_DEFAULT AlfaPipelineProps = AlfaPipelineProps{

	StackProps: environment.StackProps_DEFAULT,

	CfnConnectionProps: awscodestarconnections.CfnConnectionProps{
		ConnectionName: jsii.String("GithubConnection"),
		ProviderType:   jsii.String("GitHub"),
	},

	CodeBuildStepProps: pipelines.CodeBuildStepProps{
		Commands: jsii.Strings("npm install -g aws-cdk", "goenv install 1.19.8", "goenv local 1.19.8", "go get",
			"cd lambda && ./compile.sh handler echolambda.go; cd ..", "cdk synth"),
		Env: &map[string]*string{
			"CDK_DEFAULT_REGION":  environment.StackProps_DEFAULT.Env.Region,
			"CDK_DEFAULT_ACCOUNT": environment.StackProps_DEFAULT.Env.Account,
			"CDK_DEV_REGION":      environment.StackProps_DEV.Env.Region,
			"CDK_DEV_ACCOUNT":     environment.StackProps_DEV.Env.Account,
		},
	},

	CodePipelineProps: pipelines.CodePipelineProps{
		PipelineName:     jsii.String("EchoAppAlfa-Pipeline-default"),
		CrossAccountKeys: jsii.Bool(true),
	},

	EchoAppAlfaProps_1ENV: stages.EchoAppAlfaProps_DEFAULT,
	EchoAppAlfaProps_2ENV: stages.EchoAppAlfaProps_DEV,

	AddStageOpts: pipelines.AddStageOpts{
		Pre: &[]pipelines.Step{
			pipelines.NewManualApprovalStep(jsii.String("PromoteComponentToProduction"), &pipelines.ManualApprovalStepProps{
				Comment: jsii.String("LAST CHECK BEFORE PRODUCTION"),
			}),
		},
	},

	AlfaPipelineIds: AlfaPipelineIds{
		AlfaPipeline_Id:  "AlfaPipeline-default",
		CfnConnection_Id: "CodestarConnectionToGithub",

		CodePipelineSource_Connection_Id:        "orionverso/EchoApp_mock",
		CodePipelineSource_Connection_branch_Id: "dev",
		CodeBuildStep_Id:                        "SynthStep",
		CodePipeline_Id:                         "EchoAppAlfa-Pipeline",
		EchoAppAlfa_Id:                          "DeployStageOf-EchoAppAlfa",
	},
}

var AlfaPipeline_DEV AlfaPipelineProps = AlfaPipelineProps{

	StackProps: environment.StackProps_DEV,

	CfnConnectionProps: awscodestarconnections.CfnConnectionProps{
		ConnectionName: jsii.String("GithubConnection"),
		ProviderType:   jsii.String("GitHub"),
	},

	CodeBuildStepProps: pipelines.CodeBuildStepProps{
		Commands: jsii.Strings("npm install -g aws-cdk", "goenv install 1.19.8", "goenv local 1.19.8", "go get",
			"cd lambda && ./compile.sh handler echolambda.go; cd ..", "cdk synth"),
		Env: &map[string]*string{
			"CDK_DEV_REGION":   environment.StackProps_DEV.Env.Region,
			"CDK_DEV_ACCOUNT":  environment.StackProps_DEV.Env.Account,
			"CDK_PROD_REGION":  environment.StackProps_PROD.Env.Region,
			"CDK_PROD_ACCOUNT": environment.StackProps_PROD.Env.Account,
		},
	},

	CodePipelineProps: pipelines.CodePipelineProps{
		PipelineName:     jsii.String("EchoAppAlfa-Pipeline-dev"),
		CrossAccountKeys: jsii.Bool(true),
	},

	EchoAppAlfaProps_1ENV: stages.EchoAppAlfaProps_DEV,
	EchoAppAlfaProps_2ENV: stages.EchoAppAlfaProps_PROD,

	AddStageOpts: pipelines.AddStageOpts{
		Pre: &[]pipelines.Step{
			pipelines.NewManualApprovalStep(jsii.String("PromoteComponentToProduction"), &pipelines.ManualApprovalStepProps{
				Comment: jsii.String("LAST CHECK BEFORE PRODUCTION"),
			}),
		},
	},

	AlfaPipelineIds: AlfaPipelineIds{
		AlfaPipeline_Id:                         "AlfaPipeline-dev",
		CfnConnection_Id:                        "CodestarConnectionToGithub",
		CodePipelineSource_Connection_Id:        "orionverso/EchoApp_mock",
		CodePipelineSource_Connection_branch_Id: "dev",
		CodeBuildStep_Id:                        "SynthStep",
		CodePipeline_Id:                         "EchoAppAlfa-Pipeline",
		EchoAppAlfa_Id:                          "DeployStageOf-EchoAppAlfa",
	},
}
