package component

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type ComponentProps struct {
	awscdk.StackProps
}

type Component interface {
	NewComponentStack(constructs.Construct, *string, *ComponentProps) awscdk.Stack
	PlugComponent() Component
}
