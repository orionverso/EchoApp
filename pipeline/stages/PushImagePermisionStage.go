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
	RoleWillAssumeProps            component.RoleWillAssumeProps
	NextDeployPreparationProps_Ids
}

type nextDeployPreparation struct {
	awscdk.Stage
	assumePushImagePerm component.RolePushImageCrossAccount
	roleWillAssume      component.RoleWillAssume
}

func (as nextDeployPreparation) NextDeployPreparationStage() awscdk.Stage {
	return as.Stage
}

func (as nextDeployPreparation) RolePushImageCrossAccount() component.RolePushImageCrossAccount {
	return as.assumePushImagePerm
}

func (as nextDeployPreparation) RoleWillAssume() component.RoleWillAssume {
	return as.roleWillAssume
}

type NextDeployPreparation interface {
	NextDeployPreparationStage() awscdk.Stage
	RolePushImageCrossAccount() component.RolePushImageCrossAccount
	RoleWillAssume() component.RoleWillAssume
}

func NewNextDeployPreparation(scope constructs.Construct, id *string, props *NextDeployPreparationProps) nextDeployPreparation {

	var sprops NextDeployPreparationProps = NextDeployPreparationProps_DEV_CROSS
	var sid NextDeployPreparationProps_Ids = sprops.NextDeployPreparationProps_Ids

	if props != nil {
		sprops = *props
		sid = sprops.NextDeployPreparationProps_Ids
	}

	if id != nil {
		sid.Stage_Id = *id
	}

	stage := awscdk.NewStage(scope, jsii.String(sid.Stage_Id), &sprops.StageProps)

	cpt1 := component.NewRolePushImageCrossAccount(stage, nil, &sprops.RolePushImageCrossAccountProps)

	cpt2 := component.NewRoleWillAssume(stage, nil, &sprops.RoleWillAssumeProps)

	cpt1.RolePushImageCrossAccountRole().GrantAssumeRole(cpt2.RoleWillAssumeRole())

	return nextDeployPreparation{stage, cpt1, cpt2}

}

var NextDeployPreparationProps_DEV_CROSS NextDeployPreparationProps = NextDeployPreparationProps{
	RolePushImageCrossAccountProps: component.RolePushImageCrossAccountProps_DEV_CROSS,
	RoleWillAssumeProps:            component.RoleWillAssumeProps_DEV,
	StageProps:                     environment.Stage_PROD,
	NextDeployPreparationProps_Ids: NextDeployPreparationProps_Ids{
		Stage_Id:                 "PreparationStageToDeployNextEnviroment",
		NextDeployPreparation_Id: "AssumePushImageStage",
	},
}
