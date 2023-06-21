package writer

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type WriterFargateIds struct {
	WriterFargate_Id                         string
	ApplicationLoadBalancedFargateService_Id string
	Tag_Id                                   string
}

type WriterFargateProps struct {
	//InsideProps
	ApplicationLoadBalancedFargateServiceProps awsecspatterns.ApplicationLoadBalancedFargateServiceProps
	//Import
	Repo awsecr.Repository
	//Identifiers
	WriterFargateIds
}

type writerFargate struct {
	constructs.Construct
	repository                            awsecr.Repository
	applicationLoadBalancedFargateService awsecspatterns.ApplicationLoadBalancedFargateService
}

func (wf writerFargate) Repository() awsecr.Repository {
	return wf.repository
}

func (wf writerFargate) FargateService() awsecspatterns.ApplicationLoadBalancedFargateService {
	return wf.applicationLoadBalancedFargateService
}

func (wf writerFargate) Image() awsecs.EcrImage {
	return awsecs.AssetImage_FromEcrRepository(wf.repository, jsii.String("latest"))
}

type WriterFargate interface {
	constructs.Construct
	Repository() awsecr.Repository
	FargateService() awsecspatterns.ApplicationLoadBalancedFargateService
	Image() awsecs.EcrImage
}

func NewWriterFargate(scope constructs.Construct, id *string, props *WriterFargateProps) WriterFargate {
	//nil-case
	var sprops WriterFargateProps = WriterFargateProps_DEFAULT
	var sid WriterFargateIds = sprops.WriterFargateIds

	if props != nil {
		sprops = *props
		sid = sprops.WriterFargateIds
	}

	if id != nil {
		sid.WriterFargate_Id = *id
	}
	this := constructs.NewConstruct(scope, jsii.String(sid.WriterFargate_Id))

	image := awsecs.AssetImage_FromEcrRepository(sprops.Repo, jsii.String("latest"))

	TaskImageOptions := &awsecspatterns.ApplicationLoadBalancedTaskImageOptions{}
	TaskImageOptions.Image = image
	sprops.ApplicationLoadBalancedFargateServiceProps.TaskImageOptions = TaskImageOptions

	var cluster awsecs.Cluster
	sprops.ApplicationLoadBalancedFargateServiceProps.Cluster = cluster

	fargateservice := awsecspatterns.NewApplicationLoadBalancedFargateService(this,
		jsii.String(sid.ApplicationLoadBalancedFargateService_Id), &sprops.ApplicationLoadBalancedFargateServiceProps)

	return writerFargate{this, sprops.Repo, fargateservice}
}

// CONGIFURATIONS

var WriterFargateProps_DEFAULT WriterFargateProps = WriterFargateProps{
	ApplicationLoadBalancedFargateServiceProps: awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
		MemoryLimitMiB:   jsii.Number(1024),
		DesiredCount:     jsii.Number(1),
		Cpu:              jsii.Number(512),
		LoadBalancerName: jsii.String("echoapp-alb-default"),
	},

	WriterFargateIds: WriterFargateIds{
		WriterFargate_Id:                         "WriterFargate-default",
		ApplicationLoadBalancedFargateService_Id: "GoWebService-default",
		Tag_Id:                                   "latest",
	},
}

var WriterFargateProps_DEV WriterFargateProps = WriterFargateProps{
	ApplicationLoadBalancedFargateServiceProps: awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
		MemoryLimitMiB:   jsii.Number(1024),
		DesiredCount:     jsii.Number(1),
		Cpu:              jsii.Number(512),
		LoadBalancerName: jsii.String("echoapp-alb-dev"),
	},
	WriterFargateIds: WriterFargateIds{
		WriterFargate_Id:                         "WriterFargate-dev",
		ApplicationLoadBalancedFargateService_Id: "GoWebService-dev",
		Tag_Id:                                   "latest",
	},
}

var WriterFargateProps_PROD WriterFargateProps = WriterFargateProps{
	ApplicationLoadBalancedFargateServiceProps: awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
		MemoryLimitMiB:   jsii.Number(1024),
		DesiredCount:     jsii.Number(1),
		Cpu:              jsii.Number(512),
		LoadBalancerName: jsii.String("echoapp-alb-prod"),
	},
	WriterFargateIds: WriterFargateIds{
		WriterFargate_Id:                         "WriterFargate-prod",
		ApplicationLoadBalancedFargateService_Id: "GoWebService-prod",
		Tag_Id:                                   "latest",
	},
}
