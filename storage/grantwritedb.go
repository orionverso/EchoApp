package storage

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

func grantwriteLambda(tbl awsdynamodb.Table, lmb awslambda.Function) {
	tbl.GrantWriteData(lmb)
}
