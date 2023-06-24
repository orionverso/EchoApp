package stages

import (
	"writer_storage_app/component"
	"writer_storage_app/environment"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type NextDeployPreparationProps_Ids struct {
	Stage_Id                 string
	NextDeployPreparation_Id string
}

type NextDeployPreparationProps struct {
	StageProps                     awscdk.StageProps
	RolePushImageCrossAccountProps component.RolePushImageCrossAccountProps
	NextDeployPreparationProps_Ids
}

type nextDeployPreparation struct {
	awscdk.Stage
	assumePushImagePerm component.RolePushImageCrossAccount
}

func (as nextDeployPreparation) NextDeployPreparationStage() awscdk.Stage {
	return as.Stage
}

func (as nextDeployPreparation) RolePushImageCrossAccount() component.RolePushImageCrossAccount {
	return as.assumePushImagePerm
}

type NextDeployPreparation interface {
	NextDeployPreparationStage() awscdk.Stage
	RolePushImageCrossAccount() component.RolePushImageCrossAccount
}

func NewNextDeployPreparation(scope constructs.Construct, id *string, props *NextDeployPreparationProps) nextDeployPreparation {

	var sprops NextDeployPreparationProps = NextDeployPreparationProps_DEV
	var sid NextDeployPreparationProps_Ids = sprops.NextDeployPreparationProps_Ids

	if props != nil {
		sprops = *props
		sid = sprops.NextDeployPreparationProps_Ids
	}

	if id != nil {
		sid.NextDeployPreparation_Id = *id
	}

	stage := awscdk.NewStage(scope, jsii.String(sid.Stage_Id), &sprops.StageProps)

	cpt := component.NewRolePushImageCrossAccount(stage, nil, &sprops.RolePushImageCrossAccountProps)

	return nextDeployPreparation{stage, cpt}

}

var NextDeployPreparationProps_DEV NextDeployPreparationProps = NextDeployPreparationProps{
	RolePushImageCrossAccountProps: component.RolePushImageCrossAccountProps_DEV,
	StageProps:                     environment.Stage_DEV,
	NextDeployPreparationProps_Ids: NextDeployPreparationProps_Ids{
		Stage_Id:                 "AllowFirstEnvPushImageToSecondEnv",
		NextDeployPreparation_Id: "AssumePushImageStage",
	},
}
