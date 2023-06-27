package main

import (
	"writer_storage_app/component"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	component.NewRepo(app, nil, &component.RepoProps_DEV)

	app.Synth(nil)
}

// CONFIGURATIONS
var AppProps_DEFAULT awscdk.AppProps = awscdk.AppProps{}

var AppProps_DEV awscdk.AppProps = awscdk.AppProps{}

var AppProps_PROD awscdk.AppProps = awscdk.AppProps{}
