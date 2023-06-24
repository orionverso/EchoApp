package component

import (
	"writer_storage_app/environment"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type RolePushImageCrossAccountPropsIds struct {
	RolePushImageCrossAccount_Id      string
	Role_Id                           string
	RolePushImageCrossAccountStack_Id string
}

type RolePushImageCrossAccountProps struct {
	RoleProps  awsiam.RoleProps
	StackProps awscdk.StackProps
	RolePushImageCrossAccountPropsIds
}

type rolePushImageCrossAccount struct {
	awscdk.Stack
	role awsiam.Role
}

func (as rolePushImageCrossAccount) RolePushImageCrossAccountStack() awscdk.Stack {
	return as.Stack
}

func (as rolePushImageCrossAccount) RolePushImageCrossAccountRole() awsiam.IRole {
	return as.role
}

type RolePushImageCrossAccount interface {
	RolePushImageCrossAccountStack() awscdk.Stack
	RolePushImageCrossAccountRole() awsiam.IRole
}

func NewRolePushImageCrossAccount(scope constructs.Construct, id *string, props *RolePushImageCrossAccountProps) RolePushImageCrossAccount {

	var sprops RolePushImageCrossAccountProps = RolePushImageCrossAccountProps_DEV
	var sid RolePushImageCrossAccountPropsIds = sprops.RolePushImageCrossAccountPropsIds

	if props != nil {
		sprops = *props
		sid = sprops.RolePushImageCrossAccountPropsIds
	}

	if id != nil {
		sid.RolePushImageCrossAccount_Id = *id
	}

	stack := awscdk.NewStack(scope, jsii.String(sid.RolePushImageCrossAccountStack_Id), &sprops.StackProps)

	rl := awsiam.NewRole(stack, jsii.String(sid.Role_Id), &sprops.RoleProps)

	return rolePushImageCrossAccount{stack, rl}

}

//CONFIGURATIONS

var RolePushImageCrossAccountProps_DEV RolePushImageCrossAccountProps = RolePushImageCrossAccountProps{
	RoleProps: awsiam.RoleProps{
		AssumedBy: awsiam.NewAccountPrincipal(environment.StackProps_DEV.Env.Account),
		InlinePolicies: &map[string]awsiam.PolicyDocument{
			"AllowDevAccountPushEcr": awsiam.NewPolicyDocument(&awsiam.PolicyDocumentProps{
				Statements: &[]awsiam.PolicyStatement{
					PushImagePolicy(&PushImagePolicyProps_DEV),
				},
			}),
		},
	},

	StackProps: environment.StackProps_PROD, //Because it's cross account role

	RolePushImageCrossAccountPropsIds: RolePushImageCrossAccountPropsIds{
		RolePushImageCrossAccount_Id:      "RolePushImageCrossAccountToAssumeFromCodeBuildinFirstENV-dev",
		Role_Id:                           "AssumePushImageToAnotherAccountRole-dev",
		RolePushImageCrossAccountStack_Id: "AllowCodeBuildInFirstEnvPushImage",
	},
}

type PushImagePolicyProps struct {
	PolicyStatementProps awsiam.PolicyStatementProps
}

func PushImagePolicy(props *PushImagePolicyProps) awsiam.PolicyStatement {

	var sprops PushImagePolicyProps = PushImagePolicyProps_DEV

	if props != nil {
		sprops = *props
	}

	pl := awsiam.NewPolicyStatement(&sprops.PolicyStatementProps)

	return pl
}

var PushImagePolicyProps_DEV PushImagePolicyProps = PushImagePolicyProps{

	PolicyStatementProps: awsiam.PolicyStatementProps{
		Effect:    awsiam.Effect_ALLOW,
		Resources: jsii.Strings("*"),
		Actions: jsii.Strings(
			"ecr:GetAuthorizationToken",
			"ecr:BatchCheckLayerAvailability",
			"ecr:GetDownloadUrlForLayer",
			"ecr:GetRepositoryPolicy",
			"ecr:DescribeRepositories",
			"ecr:ListImages",
			"ecr:DescribeImages",
			"ecr:BatchGetImage",
			"ecr:InitiateLayerUpload",
			"ecr:UploadLayerPart",
			"ecr:CompleteLayerUpload",
			"ecr:PutImage",
		),
	},
}
