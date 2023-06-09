package main

import (
	"os"
	"writer_storage_app/component"
	"writer_storage_app/pipeline"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	//This is the model to deploy other stacks
	var cpt component.ApiLambdaDBComponent

	pipeline.NewPipelineStack(app, jsii.String("ComponentPipelineDev"), &pipeline.PipelineStackProps{
		DevStackProps:  awscdk.StackProps{Env: DevEnv()},
		ProdStackProps: awscdk.StackProps{Env: ProdEnv()},
		Cpt:            cpt.PlugComponent(),
	})

	app.Synth(nil)
}

func ProdEnv() *awscdk.Environment {
	region := os.Getenv("CDK_PROD_REGION")
	account := os.Getenv("CDK_PROD_ACCOUNT")

	if account != "" {
		return &awscdk.Environment{
			Region:  jsii.String(region),
			Account: jsii.String(account),
		}
	}
	return DefaultEnv()
}

func DevEnv() *awscdk.Environment {
	region := os.Getenv("CDK_DEV_REGION")
	account := os.Getenv("CDK_DEV_ACCOUNT")
	if account != "" && region != "" {
		return &awscdk.Environment{
			Region:  jsii.String(region),
			Account: jsii.String(account),
		}
	}
	return DefaultEnv()
}

func DefaultEnv() *awscdk.Environment {
	region := os.Getenv("CDK_DEFAULT_REGION")
	account := os.Getenv("CDK_DEFAULT_ACCOUNT")
	return &awscdk.Environment{
		Region:  jsii.String(region),
		Account: jsii.String(account),
	}
}
