package component

import (
	"fmt"
	"log"
	"writer_storage_app/environment"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type RoleProps struct {
	PolicyStatementProps awsiam.PolicyStatementProps
	PolicyDocumentProps  awsiam.PolicyDocumentProps
	awsiam.RoleProps
}

func (rl *RoleProps) AddUsefulDescription(arn *string) {
	rl.Description = jsii.String(fmt.Sprintf("This is a role offered by the %s to allow future interactions.The another role must create a role to assume this one. For example, CodeBuild want to use this role", *arn))
}

func (rl *RoleProps) AddSpecificResourceToStatement(arn *string) {
	*rl.PolicyStatementProps.Resources = append(*rl.PolicyStatementProps.Resources, arn)
}

func (rl *RoleProps) AddRoleName(name *string) {
	rl.RoleName = jsii.String(fmt.Sprintf("Allow_External_Interaction_With_%s", *name))
}

func (rl *RoleProps) AddStatementToPolicyDocument() {
	sts := &rl.PolicyStatementProps
	*rl.PolicyDocumentProps.Statements = append(*rl.PolicyDocumentProps.Statements, awsiam.NewPolicyStatement(sts))
}

func (rl *RoleProps) AddDocumentToInlinePolicies() {
	doc := &rl.PolicyDocumentProps
	var inline map[string]awsiam.PolicyDocument = *rl.InlinePolicies
	inline["Allow_Interaction"] = awsiam.NewPolicyDocument(doc)
}

type WithExternalInteractionIds struct {
	Role_Id string
}

type WithExternalInteractionProps struct {
	RoleProps RoleProps
	//Identifiers
	WithExternalInteractionIds
}

type withExternalInteraction struct {
	scope constructs.Construct
	role  awsiam.Role
}

func (wt withExternalInteraction) WithExternalInteractionRole() awsiam.Role {
	return wt.role
}

func (wt withExternalInteraction) WithExternalInteractionScope() constructs.Construct {
	return wt.scope
}

type WithExternalInteraction interface {
	WithExternalInteractionRole() awsiam.Role
	WithExternalInteractionScope() constructs.Construct
}

type Interactable interface {
	Arn() *string
	Name() *string
}

func NewWithExternalInteraction(scope constructs.Construct, id *string, props *WithExternalInteractionProps, ext Interactable) WithExternalInteraction {

	if props == nil {
		log.Panic("NewWithExternalInteraction is component extension. it needs an implementation of its interface")
	}

	var sprops WithExternalInteractionProps = *props
	var sid WithExternalInteractionIds = sprops.WithExternalInteractionIds

	if id != nil {
		sid.Role_Id = *id
	}

	sprops.RoleProps.AddUsefulDescription(ext.Arn())

	sprops.RoleProps.AddRoleName(ext.Name())

	sprops.RoleProps.AddSpecificResourceToStatement(ext.Arn())

	sprops.RoleProps.AddStatementToPolicyDocument()

	sprops.RoleProps.AddDocumentToInlinePolicies()

	rl := awsiam.NewRole(scope, &sprops.WithExternalInteractionIds.Role_Id, &sprops.RoleProps.RoleProps)

	return withExternalInteraction{scope, rl}
}

// CONFIGURATIONS
var WithExternalInteractionProps_REPO_DEV WithExternalInteractionProps = WithExternalInteractionProps{
	RoleProps: RoleProps{

		PolicyStatementProps: awsiam.PolicyStatementProps{
			Effect:    awsiam.Effect_ALLOW,
			Resources: &[]*string{}, //At runtime
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

		PolicyDocumentProps: awsiam.PolicyDocumentProps{
			Statements: &[]awsiam.PolicyStatement{}, //At runtime
		},

		RoleProps: awsiam.RoleProps{
			AssumedBy:      awsiam.NewAccountPrincipal(environment.StackProps_DEV.Env.Account),
			InlinePolicies: &map[string]awsiam.PolicyDocument{}, //At runtime
		},
	},

	WithExternalInteractionIds: WithExternalInteractionIds{
		Role_Id: "Allow_External_Interaction_With",
	},
}

var WithExternalInteractionProps_S3Storage_DEV WithExternalInteractionProps = WithExternalInteractionProps{
	RoleProps: RoleProps{

		PolicyStatementProps: awsiam.PolicyStatementProps{
			Effect:    awsiam.Effect_ALLOW,
			Resources: &[]*string{}, //At runtime
			Actions: jsii.Strings(
				"s3:PutObject",
				"s3:ListBucket",
			),
		},

		PolicyDocumentProps: awsiam.PolicyDocumentProps{
			Statements: &[]awsiam.PolicyStatement{}, //At runtime
		},

		RoleProps: awsiam.RoleProps{
			AssumedBy:      awsiam.NewAccountPrincipal(environment.StackProps_DEV.Env.Account),
			InlinePolicies: &map[string]awsiam.PolicyDocument{}, //At runtime
		},
	},

	WithExternalInteractionIds: WithExternalInteractionIds{
		Role_Id: "Allow_External_Interaction",
	},
}
