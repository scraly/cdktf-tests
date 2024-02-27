package main

import (
	ovh "cdk.tf/go/stack/generated/ovh/ovh/provider"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// The code that defines your stack goes here

	ovh.NewOvhProvider(stack, jsii.String("ovh"), &ovh.OvhProviderConfig{
		Endpoint: jsii.String("ovh-eu"), //required
		//ApplicationKey:    jsii.String(""),       //required OVH_APPLICATION_KEY
		//ApplicationSecret: jsii.String(""),       //required OVH_APPLICATION_SECRET
		//ConsumerKey:       jsii.String(""),       //required OVH_CONSUMER_KEY
	})

	//CloudProjectKube

	//cloudproject.NewKube()

	// Deploy a new Kubernetes cluster
	//myKube, err := cloudproject.NewKube(ctx, "my_desired_cluster", &cloudproject.KubeArgs{
	//	ServiceName: pulumi.String(serviceName),
	//	Name:        pulumi.String("my_desired_cluster"),
	//	Region:      pulumi.String("GRA5"),
	//})

	//instance := ec2.NewInstance(stack, jsii.String("cdktfgo"), &ec2.InstanceConfig{
	//    Ami:          jsii.String("ami-04c921614424b07cd"),
	//    InstanceType: jsii.String("t2.micro"),
	//})

	//cdktf.NewTerraformOutput(stack, jsii.String("public_ip"), &cdktf.TerraformOutputConfig{
	//    Value: instance.PublicIp(),
	//})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "learn-cdktf-go")

	app.Synth()
}
