package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	newPlatformStack(app, "platform-stack", awscdk.StackProps{
		StackName: jsii.String("platform-stack"),
	})

	app.Synth(nil)
}
