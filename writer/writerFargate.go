package writer

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type WriterFargateProps struct {
	RepositoryProps                            awsecr.RepositoryProps
	ApplicationLoadBalancedFargateServiceProps awsecspatterns.ApplicationLoadBalancedFargateServiceProps
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
	//Set WriterFargateProps_DEV to Default (nil-case)
	var sprops WriterFargateProps = WriterFargateProps_DEV
	if props != nil {
		sprops = *props
	}

	this := constructs.NewConstruct(scope, id)

	repo := awsecr.NewRepository(this, jsii.String("DockerEcrRepository"), &sprops.RepositoryProps)

	image := awsecs.AssetImage_FromEcrRepository(repo, jsii.String("latest"))
	//In any setting adjust this image
	sprops.ApplicationLoadBalancedFargateServiceProps.TaskImageOptions.Image = image

	fargateservice := awsecspatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("GoWebService"), &sprops.ApplicationLoadBalancedFargateServiceProps)

	return writerFargate{this, repo, fargateservice}
}

var WriterFargateProps_DEV WriterFargateProps = WriterFargateProps{
	RepositoryProps: awsecr.RepositoryProps{
		RepositoryName:   jsii.String("writer-repo-app-dev"),
		RemovalPolicy:    awscdk.RemovalPolicy_DESTROY,
		AutoDeleteImages: jsii.Bool(true),
	},
	ApplicationLoadBalancedFargateServiceProps: awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
		Cluster:          nil,
		MemoryLimitMiB:   jsii.Number(1024),
		DesiredCount:     jsii.Number(1),
		Cpu:              jsii.Number(512),
		LoadBalancerName: jsii.String("EchoApp-alb-dev-"),
	},
}

var WriterFargateProps_PROD WriterFargateProps = WriterFargateProps{
	RepositoryProps: awsecr.RepositoryProps{
		RepositoryName: jsii.String("writer-repo-app-prod"),
	},
	ApplicationLoadBalancedFargateServiceProps: awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
		Cluster:          nil,
		MemoryLimitMiB:   jsii.Number(1024),
		DesiredCount:     jsii.Number(1),
		Cpu:              jsii.Number(512),
		LoadBalancerName: jsii.String("EchoApp-alb-prod-"),
	},
}
