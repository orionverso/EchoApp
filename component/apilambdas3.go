package component

import (
	"writer_storage_app/storage"
	"writer_storage_app/writer"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ApiLambdaS3Component struct {
}

type WriterStorageAppStackApiLambdaS3Props struct {
	awscdk.StackProps
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

func (cpt ApiLambdaS3Component) NewComponentStack(scope constructs.Construct, id *string, props awscdk.StackProps) awscdk.Stack {
	return NewWriterStorageAppStackApiLambdaS3(scope, id, &WriterStorageAppStackApiLambdaS3Props{props})
}

func (cpt ApiLambdaS3Component) PlugComponent() Component {
	var component Component
	component = cpt
	return component
}
