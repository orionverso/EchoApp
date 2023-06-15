package pipeline

import (
	"writer_storage_app/component"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type EchoAppProps struct {
	StageProps awscdk.StageProps
}

type echoApp struct {
	awscdk.Stage
	cpt component.ApiLambdaDynamoDb
}

func (ec echoApp) EchoAppComponent() component.ApiLambdaDynamoDb {
	return ec.cpt
}

type EchoApp interface {
	awscdk.Stage
	EchoAppComponent() component.ApiLambdaDynamoDb
}

func EchoAppStage(scope constructs.Construct, id *string, props *EchoAppProps) EchoApp {
	var sprops awscdk.StageProps
	if props != nil {
		sprops = props.StageProps
	}
	stage := awscdk.NewStage(scope, id, &sprops)
	// TODO: Implement component interface to plug deployable  <14-06-23, orion>
	//Finally, You can connect component with the pipeline, that is the right way.
	//...For example...
	cpt := component.NewApiLambdaDynamoDb(stage, jsii.String("EchoApp"),
		&component.ApiLambdaDynamoDbProps{})

	return echoApp{stage, cpt}
	//.................
}
