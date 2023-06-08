package component

import (
	"writer_storage_app/storage"
	"writer_storage_app/writer"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FargateDBComponent struct {
}

type WriterStorageAppStackFargateDBProps struct {
	awscdk.StackProps
}

func NewWriterStorageAppStackFargateDB(scope constructs.Construct, id *string, props *WriterStorageAppStackFargateDBProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, id, &sprops)

	wr := writer.NewWriterFargate(stack, jsii.String("TaskWriter"), &writer.WriterFargateProps{})

	storage.NewDynamoDbstorage(stack, jsii.String("DynamoDbStorage"), &storage.DynamoDbstorageProps{
		PlugGranteableWriter: wr.PlugGranteableService(),
	})
	return stack
}

func (cpt FargateDBComponent) NewComponentStack(scope constructs.Construct, id *string, props awscdk.StackProps) awscdk.Stack {
	return NewWriterStorageAppStackFargateDB(scope, id, &WriterStorageAppStackFargateDBProps{props})
}

func (cpt FargateDBComponent) PlugComponent() Component {
	var component Component
	component = cpt
	return component
}
