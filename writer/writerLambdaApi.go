package writer

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type WriterApiLambdaProps struct {
	//insert props from other constructs
}

type writerApiLambda struct {
	constructs.Construct
	writerFunc awslambda.Function
}

func (wa writerApiLambda) PlugGranteableFunc() awslambda.Function {
	return wa.writerFunc
}

type WriterFunc interface {
	constructs.Construct
	PlugGranteableFunc() awslambda.Function
	//insert useful method to Do construct
}

func NewWriterApiLambda(scope constructs.Construct, id *string, props *WriterApiLambdaProps) WriterFunc {
	//implement construct
	this := constructs.NewConstruct(scope, id)

	handler := awslambda.NewFunction(this, jsii.String("EchoLambda--"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Handler: jsii.String("handler"),
		Code:    awslambda.Code_FromAsset(jsii.String("lambda/handler.zip"), nil),
	})

	logGroup := awslogs.NewLogGroup(this, jsii.String("ApiLogGroup"), &awslogs.LogGroupProps{
		LogGroupName:  jsii.String("/aws/apigateway/MyRestApi"),
		Retention:     awslogs.RetentionDays_ONE_WEEK,
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	awsapigateway.NewLambdaRestApi(this, jsii.String("EndpointWriter"), &awsapigateway.LambdaRestApiProps{
		CloudWatchRole: jsii.Bool(true),
		Handler:        handler,
		DeployOptions: &awsapigateway.StageOptions{
			StageName:            jsii.String("test"),
			DataTraceEnabled:     jsii.Bool(true),
			LoggingLevel:         awsapigateway.MethodLoggingLevel_ERROR,
			AccessLogDestination: awsapigateway.NewLogGroupLogDestination(logGroup),
		},
	})

	return writerApiLambda{this, handler}

}
