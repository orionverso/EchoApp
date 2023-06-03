package storage

import (
	"writer_storage_app/storage/choice"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// Plug-in with other constructs
type S3storageProps struct {
	PlugWriter any
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

	bucket := awss3.NewBucket(this, jsii.String("ReceiveEchoBucket"), &awss3.BucketProps{
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	ch := choice.NewChoiceStorage(this, jsii.String("StorageChoice"), &choice.ChoiceStorageProps{
		Storage_solution: jsii.String("S3"),
		Destination:      bucket.BucketName(),
	})

	pl, ok := props.PlugWriter.(awslambda.Function)

	if ok {
		bucket.GrantWrite(pl, jsii.String("*"), jsii.Strings("*"))
		ch.GrantRead(pl)
	}

	return s3storage{this}

}
