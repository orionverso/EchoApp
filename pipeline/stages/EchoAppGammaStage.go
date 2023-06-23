package stages

import (
	"writer_storage_app/component"
	"writer_storage_app/environment"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type EchoAppGammaIds struct {
	EchoAppGamma_Id string
	FargateS3_Id    string
	Repo_Id         string
}

type EchoAppGammaProps struct {
	StageProps     awscdk.StageProps
	FargateS3Props component.FargateS3Props
	RepoProps      component.RepoProps
	//Identifiers
	EchoAppGammaIds
}

type echoAppGamma struct {
	// TODO: *
	// Stage awscdk.Stage
	awscdk.Stage
	repoStack      awscdk.Stack
	fargateS3Stack awscdk.Stack
	irepo          awsecr.IRepository
}

func (ec echoAppGamma) EchoAppGammaStage() awscdk.Stage {
	return ec.Stage
}

func (ec echoAppGamma) EchoAppGammaFargateS3ComponentStack() awscdk.Stack {
	return ec.fargateS3Stack
}

func (ec echoAppGamma) EchoAppGammaRepositoryComponentStack() awscdk.Stack {
	return ec.repoStack
}

func (ec echoAppGamma) EchoAppGammaRepository() awsecr.IRepository {
	return ec.irepo
}

type EchoAppGamma interface {
	awscdk.Stage //TODO: * eliminate
	EchoAppGammaStage() awscdk.Stage
	EchoAppGammaFargateS3ComponentStack() awscdk.Stack  //Error Prone
	EchoAppGammaRepositoryComponentStack() awscdk.Stack //Order Error
	EchoAppGammaRepository() awsecr.IRepository
}

func NewEchoAppGamma(scope constructs.Construct, id *string, props *EchoAppGammaProps) EchoAppGamma {

	var sprops EchoAppGammaProps = EchoAppGammaProps_DEFAULT
	var sid EchoAppGammaIds = sprops.EchoAppGammaIds

	if props != nil {
		sprops = *props
		sid = sprops.EchoAppGammaIds
	}

	if id != nil {
		sid.EchoAppGamma_Id = *id
	}

	stage := awscdk.NewStage(scope, jsii.String(sid.EchoAppGamma_Id), &sprops.StageProps)

	rp := component.NewRepo(stage, jsii.String(sid.Repo_Id), &sprops.RepoProps)

	sprops.FargateS3Props.WriterFargateProps.Repo = rp.Repository()

	fg := component.NewFargateS3(stage, jsii.String(sid.FargateS3_Id), &sprops.FargateS3Props)
	//TODO: * This is error prone because the altered order change everything
	return echoAppGamma{stage, rp.RepoStack(), fg.FargateS3Stack(), rp.Repository()}
}

// CONFIGURATIONS
var EchoAppGammaProps_DEFAULT EchoAppGammaProps = EchoAppGammaProps{
	StageProps:     environment.Stage_DEFAULT,
	FargateS3Props: component.FargateS3Props_DEFAULT,
	RepoProps:      component.RepoProps_DEFAULT,
	EchoAppGammaIds: EchoAppGammaIds{
		EchoAppGamma_Id: "EchoApp-GAMMA-Implementation-default",
		FargateS3_Id:    "EchoApp-FargateS3-Component-default",
		Repo_Id:         "EchoApp-Repo-Component-default",
	},
}

var EchoAppGammaProps_DEV EchoAppGammaProps = EchoAppGammaProps{
	StageProps:     environment.Stage_DEV,
	FargateS3Props: component.FargateS3Props_DEV,
	RepoProps:      component.RepoProps_DEV,
	EchoAppGammaIds: EchoAppGammaIds{
		EchoAppGamma_Id: "EchoApp-GAMMA-Implementation-dev",
		FargateS3_Id:    "EchoApp-FargateS3-Component-dev",
		Repo_Id:         "EchoApp-Repo-Component-dev",
	},
}

var EchoAppGammaProps_PROD EchoAppGammaProps = EchoAppGammaProps{
	StageProps:     environment.Stage_PROD,
	FargateS3Props: component.FargateS3Props_PROD,
	RepoProps:      component.RepoProps_PROD,
	EchoAppGammaIds: EchoAppGammaIds{
		EchoAppGamma_Id: "EchoApp-GAMMA-Implementation-prod",
		FargateS3_Id:    "EchoApp-FargateS3-Component-prod",
		Repo_Id:         "EchoApp-Repo-Component-prod",
	},
}
