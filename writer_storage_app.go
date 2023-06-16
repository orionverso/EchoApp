package main

import (
	"writer_storage_app/pipeline"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	pipeline.NewLambdaPipeline(app, jsii.String("LambdaComponentPipelineDev"), &pipeline.LambdaPipelineProps{
		awscdk.StackProps{Env: DevEnv()},
		DevEnv(),
		ProdEnv(),
	})
	/*
		var cptfargate component.FargateS3Component

		pipeline.NewFargatePipelineStack(app, jsii.String("FargateComponentPipelineDev"), &pipeline.FargatePipelineStackProps{
			ProdEnv:  DevEnv(),
			CptProps: component.ComponentProps{awscdk.StackProps{Env: DevEnv()}},
			Cpt:      cptfargate.PlugComponent(),
		})
	*/

	app.Synth(nil)
}
