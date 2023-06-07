package component

import (
	"writer_storage_app/storage"
	"writer_storage_app/writer"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type WriterStorageAppStackFargateS3Props struct {
	awscdk.StackProps
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
