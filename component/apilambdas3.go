package component

import (
	"writer_storage_app/storage"
	"writer_storage_app/writer"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ApiLambdaS3ComponentProps struct {
	awscdk.StackProps
}

type ApiLambdaS3Component struct {
}

type WriterStorageAppStackApiLambdaS3Props struct {
	ApiLambdaS3ComponentProps
}

func NewWriterStorageAppStackApiLambdaS3(scope constructs.Construct, id *string, props *WriterStorageAppStackApiLambdaS3Props) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, id, &sprops)

	wr := writer.NewWriterApiLambda(stack, jsii.String("LambdaApiWriter"), &writer.WriterApiLambdaProps{})

	storage.NewS3storage(stack, jsii.String("S3Storage"), &storage.S3storageProps{
		PlugGranteableWriter: wr.PlugGranteableFunc(),
	})

	return stack
}

func (cpt ApiLambdaS3Component) NewComponentStack(scope constructs.Construct, id *string, props *ComponentProps) awscdk.Stack {
	//trangress layers
	lbs3 := ApiLambdaS3ComponentProps{props.StackProps}
	ws := WriterStorageAppStackApiLambdaS3Props{lbs3}
	//
	return NewWriterStorageAppStackApiLambdaS3(scope, id, &ws)
}

func (cpt ApiLambdaS3Component) PlugComponent() Component {
	var component Component
	component = cpt
	return component
}
