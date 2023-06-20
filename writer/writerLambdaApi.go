package writer

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type WriterApiLambdaIds struct {
	WriterApiLambda_Id string
	Function_Id        string
	LogGroup_Id        string
	LambdaRestApi_Id   string
}

type WriterApiLambdaProps struct {
	FunctionProps      awslambda.FunctionProps
	LogGroupProps      awslogs.LogGroupProps
	LambdaRestApiProps awsapigateway.LambdaRestApiProps
	//IDENTIFIERS
	WriterApiLambdaIds
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

	var sprops WriterApiLambdaProps = WriterApiLambdaProps_DEFAULT
	var sid WriterApiLambdaIds = sprops.WriterApiLambdaIds

	if props != nil {
		sprops = *props
		sid = sprops.WriterApiLambdaIds
	}

	if id != nil {
		sid.WriterApiLambda_Id = *id //Its allow the parent construct to give a name to child construct
	}

	this := constructs.NewConstruct(scope, jsii.String(sid.WriterApiLambda_Id))

	handler := awslambda.NewFunction(this, jsii.String(sid.Function_Id), &sprops.FunctionProps)

	logGroup := awslogs.NewLogGroup(this, jsii.String(sid.LogGroup_Id), &sprops.LogGroupProps)

	sprops.LambdaRestApiProps.DeployOptions.AccessLogDestination = awsapigateway.NewLogGroupLogDestination(logGroup)
	sprops.LambdaRestApiProps.Handler = handler

	apilambdaproxy := awsapigateway.NewLambdaRestApi(this, jsii.String(sid.LambdaRestApi_Id), &sprops.LambdaRestApiProps)

	return writerApiLambda{this, handler, logGroup, apilambdaproxy}
}

// CONFIGURATIONS
var WriterApiLambdaProps_DEFAULT WriterApiLambdaProps = WriterApiLambdaProps{
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
	WriterApiLambdaIds: WriterApiLambdaIds{
		WriterApiLambda_Id: "WriterApiLambda-default",
		Function_Id:        "EchoLambda-WriterToStorage-default",
		LogGroup_Id:        "EndpointWriter-logs-default",
		LambdaRestApi_Id:   "EndpointWriter-default",
	},
}

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
	WriterApiLambdaIds: WriterApiLambdaIds{
		WriterApiLambda_Id: "WriterApiLambda-dev",
		Function_Id:        "EchoLambda-WriterToStorage-dev",
		LogGroup_Id:        "EndpointWriter-logs-dev",
		LambdaRestApi_Id:   "EndpointWriter-dev",
	},
}

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
	WriterApiLambdaIds: WriterApiLambdaIds{
		WriterApiLambda_Id: "WriterApiLambda-prod",
		Function_Id:        "EchoLambda-WriterToStorage-prod",
		LogGroup_Id:        "EndpointWriter-logs-prod",
		LambdaRestApi_Id:   "EndpointWriter-prod",
	},
}
