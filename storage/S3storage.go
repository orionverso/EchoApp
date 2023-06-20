package storage

import (
	"writer_storage_app/storage/choice"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type S3StorageIds struct {
	S3Storage_Id string
	Bucket_Id    string
	Choice_Id    string
}

type S3StorageProps struct {
	BucketProps        awss3.BucketProps
	ChoiceStorageProps choice.ChoiceStorageProps
	//import from interaction with other constructs
	RoleWriter awsiam.IRole
	//Identifiers
	S3StorageIds
}

type s3Storage struct {
	constructs.Construct
	bucket awss3.Bucket
	choice choice.ChoiceStorage
}

func (s3 s3Storage) Bucket() awss3.Bucket {
	return s3.bucket
}

func (s3 s3Storage) Choice() choice.ChoiceStorage {
	return s3.choice
}

type S3Storage interface {
	constructs.Construct
	Bucket() awss3.Bucket
	Choice() choice.ChoiceStorage
}

func NewS3Storage(scope constructs.Construct, id *string, props *S3StorageProps) S3Storage {

	var sprops S3StorageProps = S3StorageProps_DEFAULT
	var sid S3StorageIds = sprops.S3StorageIds

	if props != nil {
		sprops = *props
		sid = sprops.S3StorageIds
	}

	if id != nil {
		sid.S3Storage_Id = *id
	}

	this := constructs.NewConstruct(scope, jsii.String(sid.S3Storage_Id))

	bucket := awss3.NewBucket(this, jsii.String(sid.Bucket_Id), &sprops.BucketProps)
	sprops.ChoiceStorageProps.Storage_solution = jsii.String("S3")
	sprops.ChoiceStorageProps.Destination = bucket.BucketName()

	ch := choice.NewChoiceStorage(this, jsii.String(sid.Choice_Id), &sprops.ChoiceStorageProps)

	bucket.GrantWrite(props.RoleWriter, jsii.String("*"), jsii.Strings("*"))
	ch.Storage().GrantRead(props.RoleWriter)
	ch.Destination().GrantRead(props.RoleWriter)

	return s3Storage{this, bucket, ch}
}

// CONFIGURATIONS
var S3StorageProps_DEFAULT S3StorageProps = S3StorageProps{
	BucketProps: awss3.BucketProps{
		RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
		AutoDeleteObjects: jsii.Bool(true),
	},
	ChoiceStorageProps: choice.ChoiceStorageProps_DEV,
	S3StorageIds: S3StorageIds{
		S3Storage_Id: "RecieveIn-S3Storage-default",
		Bucket_Id:    "ReceiveEchoBucket-default",
		Choice_Id:    "StorageChoice-FromBucket-default",
	},
}

var S3StorageProps_DEV S3StorageProps = S3StorageProps{
	BucketProps: awss3.BucketProps{
		RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
		AutoDeleteObjects: jsii.Bool(true),
	},
	ChoiceStorageProps: choice.ChoiceStorageProps_DEV,
	S3StorageIds: S3StorageIds{
		S3Storage_Id: "RecieveIn-S3Storage-dev",
		Bucket_Id:    "ReceiveEchoBucket-dev",
		Choice_Id:    "StorageChoice-FromBucket-dev",
	},
}

var S3StorageProps_PROD S3StorageProps = S3StorageProps{
	BucketProps:        awss3.BucketProps{},
	ChoiceStorageProps: choice.ChoiceStorageProps_PROD,
	S3StorageIds: S3StorageIds{
		S3Storage_Id: "RecieveIn-S3Storage-prod",
		Bucket_Id:    "ReceiveEchoBucket-prod",
		Choice_Id:    "StorageChoice-FromBucket-prod",
	},
}
