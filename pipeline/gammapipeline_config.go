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

type AddedStepIds struct {
	PushImageStep_Id       string
	CheckPushImageStep_Id  string
	PromoteToProduction_Id string
}

type AddedStepProps struct {
	PushImageStepProps       pipelines.CodeBuildStepProps
	CheckPushImageStepProps  pipelines.ManualApprovalStepProps
	PromoteToProductionProps pipelines.ManualApprovalStepProps
	AddedStepIds
}

func (ad AddedStepProps) AddEnvVar(index *string, value *string, props *pipelines.CodeBuildStepProps) { //Runtime Methods
	var envvars map[string]*awscodebuild.BuildEnvironmentVariable
	envvars = *props.BuildEnvironment.EnvironmentVariables
	envvars[aws.ToString(index)] = &awscodebuild.BuildEnvironmentVariable{
		Value: value,
	}
}

type GammaPipelineIds struct {
	GammaPipeline_Id                        string
	CfnConnection_Id                        string
	CodePipelineSource_Connection_Id        string
	CodePipelineSource_Connection_branch_Id string
	CodeBuildStep_Id                        string
	CodePipeline_Id                         string
	EchoAppGamma_Id                         string
	AddedStepIds
}

type GammaPipelineProps struct {
	StackProps                   awscdk.StackProps
	CfnConnectionProps           awscodestarconnections.CfnConnectionProps
	ConnectionSourceOptions      pipelines.ConnectionSourceOptions
	CodeBuildSynthStepProps      pipelines.CodeBuildStepProps
	CodePipelineProps            pipelines.CodePipelineProps
	EchoAppGammaProps_FIRST_ENV  stages.EchoAppGammaProps
	EchoAppGammaProps_SECOND_ENV stages.EchoAppGammaProps
	NextDeployPreparationProps   stages.NextDeployPreparationProps
	AddStageOpts_FIRST_ENV       pipelines.AddStageOpts
	AddStageOpts_SECOND_ENV      pipelines.AddStageOpts
	AddStageOpts_NEXT_ENV_PREP   pipelines.AddStageOpts
	AddedStepProps               AddedStepProps
	//Identifiers
	GammaPipelineIds
}

// CONFIGURATIONS

// DEFAULT: Minimal configuration to make it work

// DEV: Development configurations

var AddedStepProps_DEV AddedStepProps = AddedStepProps{
	PushImageStepProps: pipelines.CodeBuildStepProps{
		Commands: jsii.Strings(
			"cd webserver",
			"aws ecr get-login-password --region $CDK_DEV_REGION | docker login --username AWS --password-stdin $CDK_DEV_ACCOUNT.dkr.ecr.$CDK_DEV_REGION.amazonaws.com",
			"docker build -t $REPOSITORY_NAME_DEV .",
			"docker tag $REPOSITORY_NAME_DEV:latest $CDK_DEV_ACCOUNT.dkr.ecr.$CDK_DEV_REGION.amazonaws.com/$REPOSITORY_NAME_DEV:latest",
			"docker push $CDK_DEV_ACCOUNT.dkr.ecr.$CDK_DEV_REGION.amazonaws.com/$REPOSITORY_NAME_DEV:latest",
		),
		BuildEnvironment: &awscodebuild.BuildEnvironment{
			Privileged: jsii.Bool(true), //Run Docker inside CodeBuild container
			EnvironmentVariables: &map[string]*awscodebuild.BuildEnvironmentVariable{
				"CDK_DEV_REGION": &awscodebuild.BuildEnvironmentVariable{
					Value: aws.ToString(environment.StackProps_DEV.Env.Region),
				},
				"CDK_DEV_ACCOUNT": &awscodebuild.BuildEnvironmentVariable{
					Value: aws.ToString(environment.StackProps_DEV.Env.Account),
				},
				"REPOSITORY_NAME_DEV": &awscodebuild.BuildEnvironmentVariable{
					Value: aws.ToString(component.RepoProps_DEV.RepositoryProps.RepositoryName),
				},
			},
		},
		RolePolicyStatements: &[]awsiam.PolicyStatement{
			component.PushImagePolicy(&component.PushImagePolicyProps_DEV),
		},
	},

	CheckPushImageStepProps: pipelines.ManualApprovalStepProps{
		Comment: jsii.String("You can check the image was pushing"),
	},
	PromoteToProductionProps: pipelines.ManualApprovalStepProps{
		Comment: jsii.String("Last step. This component will be deployed to production"),
	},

	AddedStepIds: AddedStepIds{
		PushImageStep_Id:       "PushImageToEcrRepoStep-dev",
		CheckPushImageStep_Id:  "CheckImageInEcrRepoStep-dev",
		PromoteToProduction_Id: "PromoteToProductionStep-dev",
	},
}

var GammaPipelineProps_DEV GammaPipelineProps = GammaPipelineProps{

	StackProps: environment.StackProps_DEV,

	CfnConnectionProps: awscodestarconnections.CfnConnectionProps{
		ConnectionName: jsii.String("GithubConnection"),
		ProviderType:   jsii.String("GitHub"),
	},

	ConnectionSourceOptions: pipelines.ConnectionSourceOptions{
		TriggerOnPush: jsii.Bool(false),
	},

	CodeBuildSynthStepProps: pipelines.CodeBuildStepProps{
		Commands: jsii.Strings("npm install -g aws-cdk", "cdk synth"),
		BuildEnvironment: &awscodebuild.BuildEnvironment{ //At synth time we need all enviroment variables.
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
			},
		},
	},

	CodePipelineProps: pipelines.CodePipelineProps{
		PipelineName:      jsii.String("EchoAppGamma-Pipeline-dev"),
		CrossAccountKeys:  jsii.Bool(true),
		CodeBuildDefaults: &pipelines.CodeBuildOptions{},
	},

	EchoAppGammaProps_FIRST_ENV:  stages.EchoAppGammaProps_DEV,
	EchoAppGammaProps_SECOND_ENV: stages.EchoAppGammaProps_PROD,
	NextDeployPreparationProps:   stages.NextDeployPreparationProps_DEV_CROSS,

	AddedStepProps: AddedStepProps_DEV,

	AddStageOpts_FIRST_ENV: pipelines.AddStageOpts{
		StackSteps: &[]*pipelines.StackSteps{}, // avoid nil desrefenrence
	},

	AddStageOpts_SECOND_ENV: pipelines.AddStageOpts{
		StackSteps: &[]*pipelines.StackSteps{}, // avoid nil desrefenrence
	},

	AddStageOpts_NEXT_ENV_PREP: pipelines.AddStageOpts{
		StackSteps: &[]*pipelines.StackSteps{}, //avoid nil desreference
	},

	GammaPipelineIds: GammaPipelineIds{
		GammaPipeline_Id:                        "GammaPipeline-dev",
		CfnConnection_Id:                        "CodestarConnectionToGithub",
		CodePipelineSource_Connection_Id:        "orionverso/EchoApp_mock",
		CodePipelineSource_Connection_branch_Id: "gamma",
		CodeBuildStep_Id:                        "SynthStep",
		CodePipeline_Id:                         "EchoAppGamma-Pipeline",
		EchoAppGamma_Id:                         "DeployStageOf-EchoAppGamma",
		AddedStepIds:                            AddedStepProps_DEV.AddedStepIds,
	},
}

//PRODUCTION

var AddedStepProps_PROD AddedStepProps = AddedStepProps{
	PushImageStepProps: pipelines.CodeBuildStepProps{
		Commands: jsii.Strings(
			"cd webserver",
			"cache=\"/tmp/creds\"",
			"aws sts assume-role --role-arn $PUSH_ROLE_ARN --role-session-name pushing-prod > $cache", //test
			"export AWS_ACCESS_KEY_ID=$(cat $cache | jq -r '.Credentials.AccessKeyId')",
			"export AWS_SECRET_ACCESS_KEY=$(cat $cache | jq -r '.Credentials.SecretAccessKey')",
			"export AWS_SESSION_TOKEN=$(cat $cache | jq -r '.Credentials.SessionToken')",
			"aws ecr get-login-password --region $CDK_PROD_REGION | docker login --username AWS --password-stdin $CDK_PROD_ACCOUNT.dkr.ecr.$CDK_PROD_REGION.amazonaws.com",
			"docker build -t $REPOSITORY_NAME_PROD .",
			"docker tag $REPOSITORY_NAME_PROD:latest $CDK_PROD_ACCOUNT.dkr.ecr.$CDK_PROD_REGION.amazonaws.com/$REPOSITORY_NAME_PROD:latest",
			"docker push $CDK_PROD_ACCOUNT.dkr.ecr.$CDK_PROD_REGION.amazonaws.com/$REPOSITORY_NAME_PROD:latest",
		),

		BuildEnvironment: &awscodebuild.BuildEnvironment{
			Privileged: jsii.Bool(true), //Run Docker inside CodeBuild container
			EnvironmentVariables: &map[string]*awscodebuild.BuildEnvironmentVariable{
				"CDK_PROD_REGION": &awscodebuild.BuildEnvironmentVariable{
					Value: aws.ToString(environment.StackProps_PROD.Env.Region),
				},

				"CDK_PROD_ACCOUNT": &awscodebuild.BuildEnvironmentVariable{
					Value: aws.ToString(environment.StackProps_PROD.Env.Account),
				},
				"REPOSITORY_NAME_PROD": &awscodebuild.BuildEnvironmentVariable{
					Value: aws.ToString(component.RepoProps_PROD.RepositoryProps.RepositoryName),
				},
				/*
					"PUSH_ROLE_ARN": &awscodebuild.BuildEnvironmentVariable{
						//Value: At runtime
					},
				*/
			},
		},
		RolePolicyStatements: &[]awsiam.PolicyStatement{},
	},

	CheckPushImageStepProps: pipelines.ManualApprovalStepProps{
		Comment: jsii.String("You can check the image was pushing"),
	},

	AddedStepIds: AddedStepIds{
		PushImageStep_Id:       "PushImageToEcrRepoStep-prod",
		CheckPushImageStep_Id:  "CheckImageInEcrRepoStep-prod",
		PromoteToProduction_Id: "PromoteToProductionStep-prod",
	},
}
