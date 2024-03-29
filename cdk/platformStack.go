package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/triggers"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func newPlatformStack(scope constructs.Construct, id string, props awscdk.StackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, &props)

	bundlingOptions := &awscdklambdagoalpha.BundlingOptions{
		GoBuildFlags: &[]*string{jsii.String(`-ldflags "-s -w"`)},
	}

	// Prepare env variables
	neonDbSecrets := awssecretsmanager.Secret_FromSecretNameV2(stack, jsii.String("smartad/neon-creds"), jsii.String("smartad/neon-creds"))
	sharedEnv := map[string]*string{
		"NEON_DB_USERNAME": neonDbSecrets.SecretValueFromJson(jsii.String("username")).UnsafeUnwrap(),
		"NEON_DB_PASSWORD": neonDbSecrets.SecretValueFromJson(jsii.String("password")).UnsafeUnwrap(),
		"NEON_DB_HOST":     neonDbSecrets.SecretValueFromJson(jsii.String("host")).UnsafeUnwrap(),
		"NEON_DB_NAME":     neonDbSecrets.SecretValueFromJson(jsii.String("database")).UnsafeUnwrap(),
	}

	apiGw := awsapigateway.NewRestApi(stack, jsii.String("platform-api-gateway"), &awsapigateway.RestApiProps{
		RestApiName: jsii.String("platform-api-gateway"),
	})
	viewsTable := awsdynamodb.Table_FromTableName(stack, jsii.String("smartad-views-table"), jsii.String("smartad-views-table"))
	newViewsTable := awsdynamodb.Table_FromTableName(stack, jsii.String("sa-views"), jsii.String("sa-views"))

	userResource := apiGw.Root().AddResource(jsii.String("users"), &awsapigateway.ResourceOptions{})
	adResource := apiGw.Root().AddResource(jsii.String("ad"), &awsapigateway.ResourceOptions{}).AddResource(jsii.String("{advertisingId}"), &awsapigateway.ResourceOptions{})
	apiResourceV1 := apiGw.Root().AddResource(jsii.String("api"), &awsapigateway.ResourceOptions{}).AddResource(jsii.String("v1"), &awsapigateway.ResourceOptions{})
	apiResourceV1Proxy := apiResourceV1.AddResource(jsii.String("{proxy+}"), &awsapigateway.ResourceOptions{})
	segmentViewsResource := adResource.AddResource(jsii.String("segment-views"), &awsapigateway.ResourceOptions{})

	demoLoginHandler := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("demo-login-handler"), &awscdklambdagoalpha.GoFunctionProps{
		FunctionName: jsii.String("demo-login-handler"),
		Architecture: awslambda.Architecture_ARM_64(),
		Bundling:     bundlingOptions,
		Entry:        jsii.String("../cmd/demo-login/*.go"),
	})

	userResource.AddMethod(jsii.String("GET"),
		awsapigateway.NewLambdaIntegration(demoLoginHandler, &awsapigateway.LambdaIntegrationOptions{Proxy: jsii.Bool(true)}),
		&awsapigateway.MethodOptions{OperationName: jsii.String("demo-login")})

	getStatsHandler := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("get-pie-chart-stats-handler"), &awscdklambdagoalpha.GoFunctionProps{
		FunctionName: jsii.String("get-pie-chart-stats-handler"),
		Architecture: awslambda.Architecture_ARM_64(),
		Bundling:     bundlingOptions,
		Entry:        jsii.String("../cmd/get-pie-chart-stats/*.go"),
	})
	viewsTable.GrantReadData(getStatsHandler)
	segmentViewsResource.AddMethod(jsii.String("GET"),
		awsapigateway.NewLambdaIntegration(getStatsHandler, &awsapigateway.LambdaIntegrationOptions{Proxy: jsii.Bool(true)}),
		&awsapigateway.MethodOptions{OperationName: jsii.String("get-pie-chart-stats")})

	// General handler monolith lambda
	generalHandler := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("general-handler"), &awscdklambdagoalpha.GoFunctionProps{
		FunctionName: jsii.String("general-handler"),
		Architecture: awslambda.Architecture_ARM_64(),
		Bundling:     bundlingOptions,
		Entry:        jsii.String("../cmd/general/*.go"),
		Environment:  &sharedEnv,
	})

	viewsTable.GrantFullAccess(generalHandler)
	newViewsTable.GrantFullAccess(generalHandler)

	apiResourceV1.AddMethod(jsii.String("ANY"),
		awsapigateway.NewLambdaIntegration(generalHandler, &awsapigateway.LambdaIntegrationOptions{Proxy: jsii.Bool(true)}),
		&awsapigateway.MethodOptions{OperationName: jsii.String("general")})

	apiResourceV1Proxy.AddMethod(jsii.String("ANY"),
		awsapigateway.NewLambdaIntegration(generalHandler, &awsapigateway.LambdaIntegrationOptions{Proxy: jsii.Bool(true)}),
		&awsapigateway.MethodOptions{OperationName: jsii.String("general")})

	migratorHandler := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("migrator-handler"), &awscdklambdagoalpha.GoFunctionProps{
		FunctionName: jsii.String("migrator-handler"),
		Architecture: awslambda.Architecture_ARM_64(),
		Bundling:     bundlingOptions,
		Entry:        jsii.String("../cmd/migrator/*.go"),
		Environment:  &sharedEnv,
	})
	triggers.NewTrigger(stack, jsii.String("smartad-migration-trigger"), &triggers.TriggerProps{
		Handler:       migratorHandler,
		ExecuteBefore: &[]constructs.Construct{generalHandler},
	})

	return stack
}
