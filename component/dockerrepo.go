package component

import (
	"writer_storage_app/environment"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type RepoIds struct {
	Repository_Id string
	Stack_Id      string
}

type RepoProps struct {
	StackProps      awscdk.StackProps
	RepositoryProps awsecr.RepositoryProps
	//Identifiers
	RepoIds
}

type repo struct {
	awscdk.Stack
	repository awsecr.Repository
}

func (rp repo) Repository() awsecr.Repository {
	return rp.repository
}

func (rp repo) RepoStack() awscdk.Stack {
	return rp.Stack
}

type Repo interface {
	Repository() awsecr.Repository
	RepoStack() awscdk.Stack
}

func NewRepo(scope constructs.Construct, id *string, props *RepoProps) Repo {

	var sprops RepoProps = RepoProps_DEFAULT
	var sid RepoIds = sprops.RepoIds

	if props != nil {
		sprops = *props
		sid = sprops.RepoIds
	}
	if id != nil {
		sid.Repository_Id = *id
	}

	stack := awscdk.NewStack(scope, jsii.String(sid.Stack_Id), &sprops.StackProps)

	ecrRepo := awsecr.NewRepository(stack, jsii.String(sid.Repository_Id), &sprops.RepositoryProps)

	return repo{stack, ecrRepo}
}

//CONFIGURATIONS

var RepoProps_DEFAULT RepoProps = RepoProps{
	StackProps: environment.StackProps_DEFAULT,
	RepositoryProps: awsecr.RepositoryProps{
		RepositoryName:   jsii.String("writer-repo-app-default"),
		RemovalPolicy:    awscdk.RemovalPolicy_DESTROY,
		AutoDeleteImages: jsii.Bool(true),
	},
	RepoIds: RepoIds{
		Repository_Id: "DockerEcrRepository-default",
		Stack_Id:      "EcrRepository-default",
	},
}

var RepoProps_DEV RepoProps = RepoProps{
	StackProps: environment.StackProps_DEV,
	RepositoryProps: awsecr.RepositoryProps{
		RepositoryName:   jsii.String("writer-repo-app-dev"),
		RemovalPolicy:    awscdk.RemovalPolicy_DESTROY,
		AutoDeleteImages: jsii.Bool(true),
	},
	RepoIds: RepoIds{
		Repository_Id: "DockerEcrRepository-dev",
		Stack_Id:      "EcrRepository-dev",
	},
}

var RepoProps_PROD RepoProps = RepoProps{
	StackProps: environment.StackProps_PROD,
	RepositoryProps: awsecr.RepositoryProps{
		RepositoryName:   jsii.String("writer-repo-app-prod"),
		RemovalPolicy:    awscdk.RemovalPolicy_DESTROY,
		AutoDeleteImages: jsii.Bool(true),
	},
	RepoIds: RepoIds{
		Repository_Id: "DockerEcrRepository-prod",
		Stack_Id:      "EcrRepository-prod",
	},
}
