package component

import (
	"writer_storage_app/storage"
	"writer_storage_app/writer"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FargateS3ComponentProps struct {
	awscdk.StackProps
}

type FargateS3Component struct {
}

type WriterStorageAppStackFargateS3Props struct {
	FargateS3ComponentProps
}

func NewWriterStorageAppStackFargateS3(scope constructs.Construct, id *string, props *WriterStorageAppStackFargateS3Props) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, id, &sprops)

	wr := writer.NewWriterFargate(stack, jsii.String("TaskWriter"), &writer.WriterFargateProps{})

	storage.NewS3storage(stack, jsii.String("S3Storage"), &storage.S3storageProps{
		PlugGranteableWriter: wr.PlugGranteableService(),
	})
	return stack
}

func (cpt FargateS3Component) NewComponentStack(scope constructs.Construct, id *string, props *ComponentProps) awscdk.Stack {
	fgs3 := FargateS3ComponentProps{props.StackProps}
	ws := WriterStorageAppStackFargateS3Props{fgs3}
	return NewWriterStorageAppStackFargateS3(scope, id, &ws)
}

func (cpt FargateS3Component) PlugComponent() Component {
	var component Component
	component = cpt
	return component
}
