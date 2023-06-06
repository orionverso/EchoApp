package main

import (
	"writer_storage_app/storage"
	"writer_storage_app/writer"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type WriterStorageAppStackProps struct {
	awscdk.StackProps
}

func NewWriterStorageAppStackApiLambdaDB(scope constructs.Construct, id string, props *WriterStorageAppStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	wr := writer.NewWriterApiLambda(stack, jsii.String("LambdaApiWriter"), &writer.WriterApiLambdaProps{})

	storage.NewDynamoDbstorage(stack, jsii.String("DynamoDbStorage"), &storage.DynamoDbstorageProps{
		PlugGranteableWriter: wr.PlugGranteableFunc(),
	})

	return stack
}

func NewWriterStorageAppStackApiLambdaS3(scope constructs.Construct, id string, props *WriterStorageAppStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	wr := writer.NewWriterApiLambda(stack, jsii.String("LambdaApiWriter"), &writer.WriterApiLambdaProps{})

	storage.NewS3storage(stack, jsii.String("S3Storage"), &storage.S3storageProps{
		PlugGranteableWriter: wr.PlugGranteableFunc(),
	})

	return stack
}

func NewWriterStorageAppStackFargateS3(scope constructs.Construct, id string, props *WriterStorageAppStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	wr := writer.NewWriterFargate(stack, jsii.String("TaskWriter"), &writer.WriterFargateProps{})

	storage.NewS3storage(stack, jsii.String("S3Storage"), &storage.S3storageProps{
		PlugGranteableWriter: wr.PlugGranteableService(),
	})
	return stack
}

func NewWriterStorageAppStackFargateDB(scope constructs.Construct, id string, props *WriterStorageAppStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	wr := writer.NewWriterFargate(stack, jsii.String("TaskWriter"), &writer.WriterFargateProps{})

	storage.NewDynamoDbstorage(stack, jsii.String("DynamoDbStorage"), &storage.DynamoDbstorageProps{
		PlugGranteableWriter: wr.PlugGranteableService(),
	})
	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	/*
		NewWriterStorageAppStackDB(app, "WriterStorageAppStack", &WriterStorageAppStackProps{
			awscdk.StackProps{
				Env: env(),
			},
		})
	*/

	NewWriterStorageAppStackFargateS3(app, "WriterStorageAppStack-Fargate-S3-", &WriterStorageAppStackProps{

		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
