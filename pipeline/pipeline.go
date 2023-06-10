package pipeline

import (
	"writer_storage_app/component"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarconnections"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type PipelineStackProps struct {
	DevStackProps  awscdk.StackProps
	ProdStackProps awscdk.StackProps
	Cpt            component.Component
}

func NewPipelineStack(scope constructs.Construct, id *string, props *PipelineStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.DevStackProps
	}

	stack := awscdk.NewStack(scope, id, &sprops)
	//Connect to GitHub
	conn := awscodestarconnections.NewCfnConnection(stack, jsii.String("CodestarConnectionToGithub"),
		&awscodestarconnections.CfnConnectionProps{
			ConnectionName: jsii.String("GithubConnection"),
			ProviderType:   jsii.String("GitHub"),
		})
	//You must accepted connection manually
	//https://docs.aws.amazon.com/codepipeline/latest/userguide/connections-github.html#connections-github-cli
	githubRepo := pipelines.CodePipelineSource_Connection(jsii.String("orionverso/EchoApp_mock"),
		jsii.String("dev"),

		&pipelines.ConnectionSourceOptions{
			ConnectionArn: conn.AttrConnectionArn(),
		})

	buildTemplate := pipelines.NewCodeBuildStep(jsii.String("SynthStep"), &pipelines.CodeBuildStepProps{
		Input: githubRepo,
		Commands: jsii.Strings("npm install -g aws-cdk", "goenv install 1.19.8", "goenv local 1.19.8", "go get",
			"cd lambda && ./compile.sh handler echolambda.go; cd ..", "cdk synth"),
		Env: &map[string]*string{
			"CDK_DEV_REGION":   sprops.Env.Region,
			"CDK_DEV_ACCOUNT":  sprops.Env.Account,
			"CDK_PROD_REGION":  props.ProdStackProps.Env.Region,
			"CDK_PROD_ACCOUNT": props.ProdStackProps.Env.Account,
		},
	})

	pipe := pipelines.NewCodePipeline(stack, jsii.String("WriterStorage-PipelineStack"), &pipelines.CodePipelineProps{
		PipelineName:     jsii.String("EchoAppPipeline"),
		Synth:            buildTemplate,
		CrossAccountKeys: jsii.Bool(true),
	})

	//Development account deploy
	deployDev := EchoAppPipelineStage(stack, jsii.String("ComponentStackDev"), &EchoAppPipelineStageProps{
		stageprops: &awscdk.StageProps{Env: sprops.Env},
		CptProps:   &awscdk.StackProps{Env: sprops.Env},
		Cpt:        props.Cpt})

	pipe.AddStage(deployDev, nil)

	//Production account deploy

	deployProd := EchoAppPipelineStage(stack, jsii.String("ComponentStackProd"), &EchoAppPipelineStageProps{
		stageprops: &awscdk.StageProps{Env: props.ProdStackProps.Env},
		CptProps:   &awscdk.StackProps{Env: props.ProdStackProps.Env},
		Cpt:        props.Cpt,
	})

	pipe.AddStage(deployProd, &pipelines.AddStageOpts{
		Pre: &[]pipelines.Step{
			pipelines.NewManualApprovalStep(jsii.String("PromoteComponentToProduction"), &pipelines.ManualApprovalStepProps{
				Comment: jsii.String("LAST CHECK BEFORE PRODUCTION"),
			}),
		},
	})

	return stack
}
