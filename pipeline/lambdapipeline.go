package pipeline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarconnections"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type LambdaPipelineProps struct {
	awscdk.StackProps
	DevelopmentEnv *awscdk.Environment
	ProductionEnv  *awscdk.Environment
}

type lambdaPipeline struct {
	awscdk.Stack
}

type LambdaPipeline interface {
	awscdk.Stack
}

func NewLambdaPipeline(scope constructs.Construct, id *string, props *LambdaPipelineProps) LambdaPipeline {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	//Define Enviroment
	sprops.Env = props.DevelopmentEnv //Be explicit!

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
			"CDK_PROD_REGION":  props.ProductionEnv.Region,
			"CDK_PROD_ACCOUNT": props.ProductionEnv.Account,
		},
	})

	pipe := pipelines.NewCodePipeline(stack, jsii.String("WriterStorage-PipelineStack"), &pipelines.CodePipelineProps{
		PipelineName:     jsii.String("EchoAppPipeline"),
		Synth:            buildTemplate,
		CrossAccountKeys: jsii.Bool(true),
	})

	//Development account deploy
	deployDev := EchoAppStage(stack, jsii.String("ComponentStackDev"), &EchoAppProps{
		StageProps: awscdk.StageProps{Env: sprops.Env},
	})
	//Now, the pipeline can access to component parameters. Component ---> Pipeline , Pipeline -x-> Component
	//endpointUrl := deployDev.EchoAppComponent().ApiLambda().LambdaRestApi().Url //For example, I can use for check endpoint
	pipe.AddStage(deployDev, nil)

	//Production account deploy

	deployProd := EchoAppStage(stack, jsii.String("ComponentStackProd"), &EchoAppProps{
		StageProps: awscdk.StageProps{Env: props.ProductionEnv},
	})

	pipe.AddStage(deployProd, &pipelines.AddStageOpts{
		Pre: &[]pipelines.Step{
			pipelines.NewManualApprovalStep(jsii.String("PromoteComponentToProduction"), &pipelines.ManualApprovalStepProps{
				Comment: jsii.String("LAST CHECK BEFORE PRODUCTION"),
			}),
		},
	})

	return lambdaPipeline{stack}
}
