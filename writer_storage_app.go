package main

import (
	pipe "writer_storage_app/pipeline"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	//component.NewApiLambdaDynamoDb(app, nil, nil)

	pipe.NewAlfaPipeline(app, nil, &pipe.AlfaPipeline_DEV)

	app.Synth(nil)
}

// CONFIGURATIONS
var AppProps_DEFAULT awscdk.AppProps = awscdk.AppProps{}

var AppProps_DEV awscdk.AppProps = awscdk.AppProps{}

var AppProps_PROD awscdk.AppProps = awscdk.AppProps{}
