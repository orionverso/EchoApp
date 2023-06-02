package writer

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// Plug-in with other constructs
type WriterTaskContainerProps struct {
	//insert props from other constructs
}

type writerTaskContainer struct {
	constructs.Construct
	writerTask awsecs.TaskDefinition
}

func (wr writerTaskContainer) PlugContainer() awsecs.TaskDefinition {
	return wr.writerTask
}

type WriterTask interface {
	constructs.Construct
	PlugContainer() awsecs.TaskDefinition
}

func NewWriterInstance(scope constructs.Construct, id *string, props *WriterTaskContainerProps) WriterTask {
	//implement construct

	this := constructs.NewConstruct(scope, id)

	repo := awsecr.NewRepository(this, jsii.String("DockerEcrRepository"), &awsecr.RepositoryProps{
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	image := awsecs.AssetImage_FromEcrRepository(repo, jsii.String("latest"))

	container := awsecs.NewContainerDefinition(this, jsii.String("GoWebServerContainer"), &awsecs.ContainerDefinitionProps{
		Image: image,
	})

	task := awsecs.NewTaskDefinition(this, jsii.String("GoWebServerTask"), &awsecs.TaskDefinitionProps{})
	task.SetDefaultContainer(container)

	awsecspatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("FargateContainerWebServer"), &awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
		DesiredCount:       jsii.Number(1),
		ListenerPort:       jsii.Number(80),
		Protocol:           awselasticloadbalancingv2.ApplicationProtocol_HTTP,
		PublicLoadBalancer: jsii.Bool(true),
		Cpu:                jsii.Number(0.5),
		TaskDefinition:     task,
	})

	return writerTaskContainer{this, task}
}
