package writer

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// Plug-in with other constructs
type WriterTaskContainerProps struct {
	//insert props from other constructs
}

type writerTask struct {
	constructs.Construct
	task awsecs.TaskDefinition
}

func (wr writerTask) PlugWriter() awsecs.TaskDefinition {
	return wr.task
}

type WriterTask interface {
	constructs.Construct
	PlugWriter() awsecs.TaskDefinition
}

func NewWriterInstance(scope constructs.Construct, id *string, props *WriterTaskContainerProps) WriterTask {
	//implement construct

	this := constructs.NewConstruct(scope, id)

	repo := awsecr.NewRepository(this, jsii.String("DockerEcrRepository"), &awsecr.RepositoryProps{
		RepositoryName: jsii.String("writer-app-repo"),
		RemovalPolicy:  awscdk.RemovalPolicy_RETAIN,
	})

	image := awsecs.AssetImage_FromEcrRepository(repo, jsii.String("latest"))
	//YOU NEED TO PUSH THE FIRST IMAGE ASAP WHEN THE REPO IS CREATED

	var cluster awsecs.Cluster

	fargatepattern := awsecspatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("GoWebService"), &awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
		Cluster:        cluster,
		MemoryLimitMiB: jsii.Number(1024),
		DesiredCount:   jsii.Number(1),
		Cpu:            jsii.Number(512),
		TaskImageOptions: &awsecspatterns.ApplicationLoadBalancedTaskImageOptions{
			Image: image,
		},
		LoadBalancerName: jsii.String("application-lb-name"),
	})

	return writerTask{this, fargatepattern.TaskDefinition()}
}
