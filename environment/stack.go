package environment

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

var StackProps_DEFAULT awscdk.StackProps = awscdk.StackProps{
	Env: DefaultEnv(),
}

var StackProps_DEV awscdk.StackProps = awscdk.StackProps{
	Env: DevEnv(),
}

var StackProps_PROD awscdk.StackProps = awscdk.StackProps{
	Env: ProdEnv(),
}

func ProdEnv() *awscdk.Environment {
	region := os.Getenv("CDK_PROD_REGION")
	account := os.Getenv("CDK_PROD_ACCOUNT")

	if account != "" {
		return &awscdk.Environment{
			Region:  jsii.String(region),
			Account: jsii.String(account),
		}
	}
	return DefaultEnv()
}

func DevEnv() *awscdk.Environment {
	region := os.Getenv("CDK_DEV_REGION")
	account := os.Getenv("CDK_DEV_ACCOUNT")
	if account != "" && region != "" {
		return &awscdk.Environment{
			Region:  jsii.String(region),
			Account: jsii.String(account),
		}
	}
	return DefaultEnv()
}

func DefaultEnv() *awscdk.Environment {
	region := os.Getenv("CDK_DEFAULT_REGION")
	account := os.Getenv("CDK_DEFAULT_ACCOUNT")
	return &awscdk.Environment{
		Region:  jsii.String(region),
		Account: jsii.String(account),
	}
}
