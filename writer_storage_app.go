package main

import (
	"writer_storage_app/component"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	//component.NewApiLambdaDynamoDb(app, nil, nil)
	sprops := component.FargateS3Props_DEV
	component.NewFargateS3(app, nil, &sprops)

	//pipeline.NewAlfaPipeline(app, nil, nil)

	app.Synth(nil)
}

// CONFIGURATIONS
var AppProps_DEFAULT awscdk.AppProps = awscdk.AppProps{}

var AppProps_DEV awscdk.AppProps = awscdk.AppProps{}

var AppProps_PROD awscdk.AppProps = awscdk.AppProps{}
