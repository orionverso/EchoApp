package storage

import (
	"writer_storage_app/storage/choice"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type S3storageProps struct {
	//import
	PlugGranteableWriter awsiam.IGrantable
}

type s3storage struct {
	constructs.Construct
	//export
}

type S3storage interface {
	constructs.Construct
}

func NewS3storage(scope constructs.Construct, id *string, props *S3storageProps) S3storage {

	this := constructs.NewConstruct(scope, id)

	bucket := awss3.NewBucket(this, jsii.String("ReceiveEchoBucket"), &awss3.BucketProps{
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	ch := choice.NewChoiceStorage(this, jsii.String("StorageChoice"), &choice.ChoiceStorageProps{
		Storage_solution: jsii.String("S3"),
		Destination:      bucket.BucketName(),
	})

	bucket.GrantWrite(props.PlugGranteableWriter, jsii.String("*"), jsii.Strings("*"))
	ch.GrantRead(props.PlugGranteableWriter)

	return s3storage{this}

}
