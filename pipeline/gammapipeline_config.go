package pipeline

import (
	"writer_storage_app/component"
	"writer_storage_app/environment"
	"writer_storage_app/pipeline/stages"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarconnections"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/aws-sdk-go-v2/aws"
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

type AddedStep struct {
	PushImageStep       pipelines.CodeBuildStep
	CheckPushImageStep  pipelines.ManualApprovalStep
	PromoteToProduction pipelines.ManualApprovalStep
}

type GammaPipelineProps struct {
	StackProps                   awscdk.StackProps
	CfnConnectionProps           awscodestarconnections.CfnConnectionProps
	ConnectionSourceOptions      pipelines.ConnectionSourceOptions
	CodeBuildStepProps           pipelines.CodeBuildStepProps
	CodePipelineProps            pipelines.CodePipelineProps
	EchoAppGammaProps_FIRST_ENV  stages.EchoAppGammaProps
	EchoAppGammaProps_SECOND_ENV stages.EchoAppGammaProps
	AddStageOpts_FIRST_ENV       pipelines.AddStageOpts
	AddStageOpts_SECOND_ENV      pipelines.AddStageOpts
	AddedStep                    AddedStep
	//Identifiers
	GammaPipelineIds
}

// CONFIGURATIONS

// DEFAULT: Minimal configuration to make it work

// DEV: Development configurations

var AddedStep_DEV AddedStep = AddedStep{
	PushImageStep: pipelines.NewCodeBuildStep(jsii.String("PushImagetoRepo"), &pipelines.CodeBuildStepProps{
		Commands: jsii.Strings(
			"cd webserver",
			"aws ecr get-login-password --region $CDK_DEV_REGION | docker login --username AWS --password-stdin $CDK_DEV_ACCOUNT.dkr.ecr.$CDK_DEV_REGION.amazonaws.com",
			"docker build -t $REPOSITORY_NAME_DEV .",
			"docker tag $REPOSITORY_NAME_DEV:latest $CDK_DEV_ACCOUNT.dkr.ecr.$CDK_DEV_REGION.amazonaws.com/$REPOSITORY_NAME_DEV:latest",
			"docker push $CDK_DEV_ACCOUNT.dkr.ecr.$CDK_DEV_REGION.amazonaws.com/$REPOSITORY_NAME_DEV:latest",
		),

		BuildEnvironment: &awscodebuild.BuildEnvironment{
			Privileged: jsii.Bool(true), //Run Docker inside CodeBuild container
		},
	}),

	CheckPushImageStep: pipelines.NewManualApprovalStep(jsii.String("CheckImagePush"), &pipelines.ManualApprovalStepProps{
		Comment: jsii.String("You can check the image was pushing"),
	}),
	PromoteToProduction: pipelines.NewManualApprovalStep(jsii.String("ProductionPromotion"), &pipelines.ManualApprovalStepProps{
		Comment: jsii.String("Last step. This component will be deployed to production"),
	}),
}

var GammaPipelineProps_DEV GammaPipelineProps = GammaPipelineProps{

	StackProps: environment.StackProps_DEV,

	CfnConnectionProps: awscodestarconnections.CfnConnectionProps{
		ConnectionName: jsii.String("GithubConnection"),
		ProviderType:   jsii.String("GitHub"),
	},

	CodeBuildStepProps: pipelines.CodeBuildStepProps{
		Commands: jsii.Strings("npm install -g aws-cdk", "cdk synth"),
	},

	CodePipelineProps: pipelines.CodePipelineProps{
		PipelineName:     jsii.String("EchoAppGamma-Pipeline-dev"),
		CrossAccountKeys: jsii.Bool(true),
		CodeBuildDefaults: &pipelines.CodeBuildOptions{
			RolePolicy: &[]awsiam.PolicyStatement{
				PushEcrPolicy(), //docker login sirve para dev account no para prod account
			},
			BuildEnvironment: &awscodebuild.BuildEnvironment{
				Privileged: jsii.Bool(true), //You need for build docker images inside codebuild project
				EnvironmentVariables: &map[string]*awscodebuild.BuildEnvironmentVariable{
					"CDK_DEV_REGION": &awscodebuild.BuildEnvironmentVariable{
						Value: aws.ToString(environment.StackProps_DEV.Env.Region),
					},
					"CDK_DEV_ACCOUNT": &awscodebuild.BuildEnvironmentVariable{
						Value: aws.ToString(environment.StackProps_DEV.Env.Account),
					},
					"CDK_PROD_REGION": &awscodebuild.BuildEnvironmentVariable{
						Value: aws.ToString(environment.StackProps_PROD.Env.Region),
					},

					"CDK_PROD_ACCOUNT": &awscodebuild.BuildEnvironmentVariable{
						Value: aws.ToString(environment.StackProps_PROD.Env.Account),
					},

					"REPOSITORY_NAME_DEV": &awscodebuild.BuildEnvironmentVariable{ //At runtime
						Value: aws.ToString(component.RepoProps_DEV.RepositoryProps.RepositoryName),
					},
					"REPOSITORY_NAME_PROD": &awscodebuild.BuildEnvironmentVariable{ //At runtime
						Value: aws.ToString(component.RepoProps_PROD.RepositoryProps.RepositoryName),
					},
				},
			},
		},
	},

	EchoAppGammaProps_FIRST_ENV:  stages.EchoAppGammaProps_DEV,
	EchoAppGammaProps_SECOND_ENV: stages.EchoAppGammaProps_PROD,

	AddedStep: AddedStep_DEV,

	AddStageOpts_FIRST_ENV: pipelines.AddStageOpts{
		StackSteps: &[]*pipelines.StackSteps{
			&pipelines.StackSteps{ //RepoStack
				//Stack: At runtime
				Post: &[]pipelines.Step{
					AddedStep_DEV.PushImageStep,
					AddedStep_DEV.CheckPushImageStep,
				},
			},
			&pipelines.StackSteps{ //FargateStack
				//Stack: At runtime
				Post: &[]pipelines.Step{
					AddedStep_DEV.PromoteToProduction,
				},
			},
		},
	},

	AddStageOpts_SECOND_ENV: pipelines.AddStageOpts{
		StackSteps: &[]*pipelines.StackSteps{
			&pipelines.StackSteps{
				//Repo Stack pass at runtime
				Post: &[]pipelines.Step{
					AddedStep_PROD.PushImageStep,
					AddedStep_PROD.CheckPushImageStep,
				},
			},
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

//PRODUCTION

var AddedStep_PROD AddedStep = AddedStep{
	PushImageStep: pipelines.NewCodeBuildStep(jsii.String("PushImagetoRepo"), &pipelines.CodeBuildStepProps{
		Commands: jsii.Strings(
			"cd webserver",
			"aws ecr get-login-password --region $CDK_PROD_REGION | docker login --username AWS --password-stdin $CDK_PROD_ACCOUNT.dkr.ecr.$CDK_PROD_REGION.amazonaws.com",
			"docker build -t $REPOSITORY_NAME_PROD .",
			"docker tag $REPOSITORY_NAME_PROD:latest $CDK_PROD_ACCOUNT.dkr.ecr.$CDK_PROD_REGION.amazonaws.com/$REPOSITORY_NAME_PROD:latest",
			"docker push $CDK_PROD_ACCOUNT.dkr.ecr.$CDK_PROD_REGION.amazonaws.com/$REPOSITORY_NAME_PROD:latest",
		),

		BuildEnvironment: &awscodebuild.BuildEnvironment{
			Privileged: jsii.Bool(true), //Run Docker inside CodeBuild container
		},
	}),

	CheckPushImageStep: pipelines.NewManualApprovalStep(jsii.String("CheckImagePush"), &pipelines.ManualApprovalStepProps{
		Comment: jsii.String("You can check the image was pushing"),
	}),
}
