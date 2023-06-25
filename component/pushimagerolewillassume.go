package component

import (
	"writer_storage_app/environment"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type RoleWillAssumePropsIds struct {
	RoleWillAssume_Id      string
	Role_Id                string
	RoleWillAssumeStack_Id string
}

type RoleWillAssumeProps struct {
	RoleProps  awsiam.RoleProps
	StackProps awscdk.StackProps
	RoleWillAssumePropsIds
}

type roleWillAssume struct {
	awscdk.Stack
	role awsiam.Role
}

func (as roleWillAssume) RoleWillAssumeStack() awscdk.Stack {
	return as.Stack
}

func (as roleWillAssume) RoleWillAssumeRole() awsiam.IRole {
	return as.role
}

type RoleWillAssume interface {
	RoleWillAssumeStack() awscdk.Stack
	RoleWillAssumeRole() awsiam.IRole
}

func NewRoleWillAssume(scope constructs.Construct, id *string, props *RoleWillAssumeProps) RoleWillAssume {

	var sprops RoleWillAssumeProps = RoleWillAssumeProps_DEV
	var sid RoleWillAssumePropsIds = sprops.RoleWillAssumePropsIds

	if props != nil {
		sprops = *props
		sid = sprops.RoleWillAssumePropsIds
	}

	if id != nil {
		sid.RoleWillAssumeStack_Id = *id //TODO: Can overwrite only construct names.
	}

	stack := awscdk.NewStack(scope, jsii.String(sid.RoleWillAssumeStack_Id), &sprops.StackProps)

	rl := awsiam.NewRole(stack, jsii.String(sid.Role_Id), &sprops.RoleProps)

	return roleWillAssume{stack, rl}

}

//CONFIGURATIONS

var RoleWillAssumeProps_DEV RoleWillAssumeProps = RoleWillAssumeProps{
	RoleProps: awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("codebuild.amazonaws.com"), &awsiam.ServicePrincipalOpts{}),
	},

	StackProps: environment.StackProps_DEV, //Because it's cross account role

	RoleWillAssumePropsIds: RoleWillAssumePropsIds{
		RoleWillAssume_Id:      "RoleWillAssumeFromCodeBuildinFirstENV",
		Role_Id:                "RoleWillAssumeToPushImageCrossAccount-dev",
		RoleWillAssumeStack_Id: "AllowCodeBuildPushCrossAccount",
	},
}
