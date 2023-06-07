package pipeline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarconnections"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type PipelineStackProps struct {
	awscdk.StackProps
}

func NewPipelineStack(scope constructs.Construct, id *string, props *PipelineStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, id, &sprops)

	conn := awscodestarconnections.NewCfnConnection(stack, jsii.String("CodestarConnectionToGithub"),
		&awscodestarconnections.CfnConnectionProps{
			ConnectionName: jsii.String("GithubConnection"),
			ProviderType:   jsii.String("GitHub"),
		})
	//You must accepted connection manually
	//https://docs.aws.amazon.com/codepipeline/latest/userguide/connections-github.html#connections-github-cli
	githubRepo := pipelines.CodePipelineSource_Connection(jsii.String("orionverso/EchoApp"),
		jsii.String("main"),
		&pipelines.ConnectionSourceOptions{
			ConnectionArn: conn.AttrConnectionArn(),
		})

	buildTemplate := pipelines.NewCodeBuildStep(jsii.String("SynthStep"), &pipelines.CodeBuildStepProps{
		Input:    githubRepo,
		Commands: jsii.Strings("npm install -g aws-cdk", "goenv install 1.19.8", "goenv local 1.19.8", "go get", "cdk synth"),
	})

	pipe := pipelines.NewCodePipeline(stack, jsii.String("WriterStorage-PipelineStack"), &pipelines.CodePipelineProps{
		PipelineName: jsii.String("EchoAppPipeline"),
		Synth:        buildTemplate,
	})

	deploy := EchoAppPipelineStage(stack, jsii.String("ApiLambdaDynamoDbComponent"), nil)

	pipe.AddStage(deploy, nil)

	return stack
}
