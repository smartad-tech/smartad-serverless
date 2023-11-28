package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
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

	apiGw := awsapigateway.NewRestApi(stack, jsii.String("platform-api-gateway"), &awsapigateway.RestApiProps{
		RestApiName: jsii.String("platform-api-gateway"),
	})
	viewsTable := awsdynamodb.Table_FromTableName(stack, jsii.String("smartad-views-table"), jsii.String("smartad-views-table"))

	userResource := apiGw.Root().AddResource(jsii.String("users"), &awsapigateway.ResourceOptions{})
	adResource := apiGw.Root().AddResource(jsii.String("ad"), &awsapigateway.ResourceOptions{}).AddResource(jsii.String("{advertisingId}"), &awsapigateway.ResourceOptions{})
	adViewsResource := adResource.AddResource(jsii.String("views"), &awsapigateway.ResourceOptions{})

	demoLoginHandler := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("demo-login-handler"), &awscdklambdagoalpha.GoFunctionProps{
		FunctionName: jsii.String("demo-login-handler"),
		Architecture: awslambda.Architecture_ARM_64(),
		Bundling:     bundlingOptions,
		Entry:        jsii.String("../cmd/demo-login/*.go"),
	})

	userResource.AddMethod(jsii.String("GET"),
		awsapigateway.NewLambdaIntegration(demoLoginHandler, &awsapigateway.LambdaIntegrationOptions{Proxy: jsii.Bool(true)}),
		&awsapigateway.MethodOptions{OperationName: jsii.String("demo-login")})

	getStatsHandler := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("get-stats-handler"), &awscdklambdagoalpha.GoFunctionProps{
		FunctionName: jsii.String("get-stats-handler"),
		Architecture: awslambda.Architecture_ARM_64(),
		Bundling:     bundlingOptions,
		Entry:        jsii.String("../cmd/get-stats/*.go"),
	})
	viewsTable.GrantReadData(getStatsHandler)
	adViewsResource.AddMethod(jsii.String("GET"),
		awsapigateway.NewLambdaIntegration(getStatsHandler, &awsapigateway.LambdaIntegrationOptions{Proxy: jsii.Bool(true)}),
		&awsapigateway.MethodOptions{OperationName: jsii.String("get-stats")})

	return stack
}
