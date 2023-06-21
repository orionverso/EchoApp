package main

import (
	pipe "writer_storage_app/pipeline"
	"writer_storage_app/writer/asset"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	//component.NewApiLambdaDynamoDb(app, nil, nil)
	assetRepo := asset.NewRepo(app, nil, &asset.RepoProps_DEV)

	sprops := pipe.GammaPipeline_DEV

	sprops.EchoAppGammaProps_1ENV.FargateS3Props.WriterFargateProps.Repo = assetRepo.Repository()
	sprops.EchoAppGammaProps_2ENV.FargateS3Props.WriterFargateProps.Repo = assetRepo.Repository()

	pipe.NewGammaPipeline(app, nil, &sprops)

	app.Synth(nil)
}

// CONFIGURATIONS
var AppProps_DEFAULT awscdk.AppProps = awscdk.AppProps{}

var AppProps_DEV awscdk.AppProps = awscdk.AppProps{}

var AppProps_PROD awscdk.AppProps = awscdk.AppProps{}
