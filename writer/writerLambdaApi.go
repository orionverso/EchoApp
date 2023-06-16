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
	FunctionProps      awslambda.FunctionProps
	LogGroupProps      awslogs.LogGroupProps
	LambdaRestApiProps awsapigateway.LambdaRestApiProps
}

type writerApiLambda struct {
	constructs.Construct
	function      awslambda.Function
	loggroup      awslogs.LogGroup
	lambdarestapi awsapigateway.LambdaRestApi
}

func (wa writerApiLambda) Function() awslambda.Function {
	return wa.function
}

func (wa writerApiLambda) LogGroup() awslogs.LogGroup {
	return wa.loggroup
}

func (wa writerApiLambda) LambdaRestApi() awsapigateway.LambdaRestApi {
	return wa.lambdarestapi
}

type WriterApiLambda interface {
	constructs.Construct
	Function() awslambda.Function
	LogGroup() awslogs.LogGroup
	LambdaRestApi() awsapigateway.LambdaRestApi
}

func NewWriterApiLambda(scope constructs.Construct, id *string, props *WriterApiLambdaProps) WriterApiLambda {
	//Set WriterApiLambdaProps_DEV to Default (nil-case)
	var sprops WriterApiLambdaProps = WriterApiLambdaProps_DEV
	if props != nil {
		sprops = *props
	}
	this := constructs.NewConstruct(scope, id)

	handler := awslambda.NewFunction(this, jsii.String("EchoLambda--"), &sprops.FunctionProps)

	logGroup := awslogs.NewLogGroup(this, jsii.String("ApiLogGroup"), &sprops.LogGroupProps)

	sprops.LambdaRestApiProps.DeployOptions.AccessLogDestination = awsapigateway.NewLogGroupLogDestination(logGroup)
	sprops.LambdaRestApiProps.Handler = handler

	apilambdaproxy := awsapigateway.NewLambdaRestApi(this, jsii.String("EndpointWriter"), &sprops.LambdaRestApiProps)

	return writerApiLambda{this, handler, logGroup, apilambdaproxy}
}

// CONFIGURATIONS
var WriterApiLambdaProps_DEV WriterApiLambdaProps = WriterApiLambdaProps{
	FunctionProps: awslambda.FunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Handler: jsii.String("handler"),
		Code:    awslambda.Code_FromAsset(jsii.String("lambda/handler.zip"), nil),
	},
	LogGroupProps: awslogs.LogGroupProps{
		LogGroupName:  jsii.String("/aws/apigateway/MyRestApi"),
		Retention:     awslogs.RetentionDays_ONE_WEEK,
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	},
	LambdaRestApiProps: awsapigateway.LambdaRestApiProps{
		CloudWatchRole: jsii.Bool(true),
		DeployOptions: &awsapigateway.StageOptions{
			StageName:        jsii.String("test"),
			DataTraceEnabled: jsii.Bool(true),
			LoggingLevel:     awsapigateway.MethodLoggingLevel_ERROR,
		},
	},
}

// TODO: implement production configurations
var WriterApiLambdaProps_PROD WriterApiLambdaProps = WriterApiLambdaProps{
	FunctionProps: awslambda.FunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Handler: jsii.String("handler"),
		Code:    awslambda.Code_FromAsset(jsii.String("lambda/handler.zip"), nil),
	},
	LogGroupProps: awslogs.LogGroupProps{
		LogGroupName:  jsii.String("/aws/apigateway/MyRestApi"),
		Retention:     awslogs.RetentionDays_ONE_WEEK,
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	},
	LambdaRestApiProps: awsapigateway.LambdaRestApiProps{
		CloudWatchRole: jsii.Bool(true),
		DeployOptions: &awsapigateway.StageOptions{
			StageName:        jsii.String("test"),
			DataTraceEnabled: jsii.Bool(true),
			LoggingLevel:     awsapigateway.MethodLoggingLevel_ERROR,
		},
	},
}
