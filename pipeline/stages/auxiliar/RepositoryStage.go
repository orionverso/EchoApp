package auxiliar

import (
	"writer_storage_app/environment"
	"writer_storage_app/writer/asset"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type EcrRepositoryIds struct {
	Stage_id string
	Repo_id  string
}

type EcrRepositoryProps struct {
	StageProps awscdk.StageProps
	RepoProps  asset.RepoProps
	//Identifiers
	EcrRepositoryIds
}

type ecrRepository struct {
	awscdk.Stage
	Repo asset.Repo
}

func (ec ecrRepository) EcrRepositoryStage() awscdk.Stage {
	return ec.Stage
}

func (ec ecrRepository) EcrRepositoryRepo() asset.Repo {
	return ec.Repo
}

type EcrRepository interface {
	EcrRepositoryStage() awscdk.Stage
	EcrRepositoryRepo() asset.Repo
}

func NewEcrRepository(scope constructs.Construct, id *string, props *EcrRepositoryProps) EcrRepository {
	var sprops EcrRepositoryProps = EcrRepositoryProps_DEFAULT
	var sid EcrRepositoryIds = sprops.EcrRepositoryIds

	if props != nil {
		sprops = *props
		sid = sprops.EcrRepositoryIds
	}

	if id != nil {
		sid.Repo_id = *id
	}

	stage := awscdk.NewStage(scope, jsii.String(sid.Stage_id), &sprops.StageProps)

	repo := asset.NewRepo(stage, jsii.String(sid.Repo_id), &sprops.RepoProps)

	return ecrRepository{stage, repo}
}

//CONFIGURATIONS

var EcrRepositoryProps_DEFAULT EcrRepositoryProps = EcrRepositoryProps{
	StageProps: environment.Stage_DEFAULT,
	RepoProps:  asset.RepoProps_DEFAULT,
	EcrRepositoryIds: EcrRepositoryIds{
		Stage_id: "EcrRepository-stage-default",
		Repo_id:  "repository-to-writer-default",
	},
}

var EcrRepositoryProps_DEV EcrRepositoryProps = EcrRepositoryProps{
	StageProps: environment.Stage_DEV,
	RepoProps:  asset.RepoProps_DEV,
	EcrRepositoryIds: EcrRepositoryIds{
		Stage_id: "EcrRepository-stage-dev",
		Repo_id:  "repository-to-writer-dev",
	},
}

var EcrRepositoryProps_PROD EcrRepositoryProps = EcrRepositoryProps{
	StageProps: environment.Stage_PROD,
	RepoProps:  asset.RepoProps_PROD,
	EcrRepositoryIds: EcrRepositoryIds{
		Stage_id: "EcrRepository-stage-prod",
		Repo_id:  "repository-to-writer-prod",
	},
}
