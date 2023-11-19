package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func newPlatformStack(scope constructs.Construct, id string, props awscdk.StackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, &props)

	bundlingOptions := &awscdklambdagoalpha.BundlingOptions{
		GoBuildFlags: &[]*string{jsii.String(`-ldflags "-s -w"`)},
	}

	awscdklambdagoalpha.NewGoFunction(stack, jsii.String("demo-login-handler"), &awscdklambdagoalpha.GoFunctionProps{
		FunctionName: jsii.String("demo-login-handler"),
		Architecture: awslambda.Architecture_ARM_64(),
		Bundling: bundlingOptions,
		Entry: jsii.String("../cmd/demo-login/*.go"), 
	})
	return stack
}
