package environment

import "github.com/aws/aws-cdk-go/awscdk/v2"

var Stage_DEFAULT awscdk.StageProps = awscdk.StageProps{
	Env: StackProps_DEFAULT.Env,
}

var Stage_DEV awscdk.StageProps = awscdk.StageProps{
	Env: StackProps_DEV.Env,
}

var Stage_PROD awscdk.StageProps = awscdk.StageProps{
	Env: StackProps_PROD.Env,
}
