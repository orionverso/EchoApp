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

	this := constructs.NewConstruct(scope, id)

	repo := awsecr.NewRepository(this, jsii.String("DockerEcrRepository"), &awsecr.RepositoryProps{
		RepositoryName: jsii.String("writer-app-repo"),
		RemovalPolicy:  awscdk.RemovalPolicy_DESTROY,
	})

	image := awsecs.AssetImage_FromEcrRepository(repo, jsii.String("latest"))
	//YOU NEED TO PUSH THE FIRST IMAGE ASAP WHEN THE REPO IS CREATED

	var cluster awsecs.Cluster

	fargateservice := awsecspatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("GoWebService"), &awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
		Cluster:        cluster,
		MemoryLimitMiB: jsii.Number(1024),
		DesiredCount:   jsii.Number(1),
		Cpu:            jsii.Number(512),
		TaskImageOptions: &awsecspatterns.ApplicationLoadBalancedTaskImageOptions{
			Image: image,
		},
		LoadBalancerName: jsii.String("application-lb-name"),
	})

	return writerFargate{this, repo, fargateservice}
}
