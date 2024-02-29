package main

import (
	"cdk.tf/go/stack/generated/kreuzwerker/docker/container"
	"cdk.tf/go/stack/generated/kreuzwerker/docker/image"
	"cdk.tf/go/stack/generated/kreuzwerker/docker/network"
	docker "cdk.tf/go/stack/generated/kreuzwerker/docker/provider"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

const (
	imageTag              = "linux-amd64"
	gophersAPIPort        = 8080
	gophersAPIWatcherPort = 8000
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// The code that defines your stack goes here

	// Initialize the Docker provider
	docker.NewDockerProvider(stack, jsii.String("docker"), &docker.DockerProviderConfig{})

	// Pull the Gophers API image
	gophersAPIImage := image.NewImage(stack, jsii.String("gophersAPIImage"), &image.ImageConfig{
		Name:        jsii.String("scraly/gophers-api:" + imageTag),
		KeepLocally: jsii.Bool(false),
	})

	// Pull the Gophers API Watcher image
	gophersAPIWatcherImage := image.NewImage(stack, jsii.String("gophersAPIWatcherImage"), &image.ImageConfig{
		Name:        jsii.String("scraly/gophers-api-watcher:" + imageTag),
		KeepLocally: jsii.Bool(false),
	})

	// Create a Docker network to allows our containers to comunicate to each others
	//TODO: xxxx
	network.NewNetwork(stack, jsii.String("network"), &network.NetworkConfig{
		Name: jsii.String("gophers"),
	})

	//ctx.Export("containerNetwork", network.Name)

	// Create the Gophers API container
	container.NewContainer(stack, jsii.String("gophersAPIContainer"), &container.ContainerConfig{
		Image: gophersAPIImage.Name(),
		Name:  jsii.String("gophers-api"),
		Ports: &[]*container.ContainerPorts{{
			Internal: jsii.Number(gophersAPIPort), External: jsii.Number(gophersAPIPort),
		}},
	})

	// Create the Gophers API Watcher container
	container.NewContainer(stack, jsii.String("gophersAPIWatcherContainer"), &container.ContainerConfig{
		Image: gophersAPIWatcherImage.Name(),
		Name:  jsii.String("gophers-api-watcher"),
		Ports: &[]*container.ContainerPorts{{
			Internal: jsii.Number(gophersAPIWatcherPort), External: jsii.Number(gophersAPIWatcherPort),
		}},
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "docker")

	app.Synth()
}
