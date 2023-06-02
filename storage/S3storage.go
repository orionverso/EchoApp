package storage

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// Plug-in with other constructs
type S3storageProps struct {
	awslambda.Function
	//insert props from other constructs
}

type s3storage struct {
	constructs.Construct
	//insert construct from other resources
}

type S3storage interface {
	constructs.Construct
	//insert useful method to Do construct
}

func NewS3storage(scope constructs.Construct, id *string, props *S3storageProps) S3storage {
	//implement construct
	this := constructs.NewConstruct(scope, id)

	bucket := awss3.NewBucket(this, jsii.String("ReceiveEchoBucket"), &awss3.BucketProps{})

	bucket.GrantWrite(props.Function, jsii.String("*"), jsii.Strings("*"))

	props.Function.AddEnvironment(jsii.String("STORAGE_SOLUTION"), jsii.String("S3"), &awslambda.EnvironmentOptions{})
	props.Function.AddEnvironment(jsii.String("DESTINATION"), bucket.BucketArn(), &awslambda.EnvironmentOptions{})

	return s3storage{this}

}
