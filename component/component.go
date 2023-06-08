package component

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type Component interface {
	NewComponentStack(constructs.Construct, *string, awscdk.StackProps) awscdk.Stack
	PlugComponent() Component
}
