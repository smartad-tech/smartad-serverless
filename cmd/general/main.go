package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/smartad-tech/smartad-serverless/cmd/general/factory"
	"github.com/smartad-tech/smartad-serverless/cmd/general/routes"
	"github.com/smartad-tech/smartad-serverless/internal/database"
)

var fiberLambda *fiberadapter.FiberLambda
var app *fiber.App

// init the Fiber Server
func init() {
	log.Printf("Fiber cold start...")
	app = fiber.New()

	// Temporary setting
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, https://app.smartad.tech",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// DI
	dynamoDbClient := database.InitDynamo()
	repositories := factory.InitRepositories(dynamoDbClient)
	controllers := factory.InitControllers(repositories)

	// Add public routes
	routes.PublicRoutes(app, controllers)

	fiberLambda = fiberadapter.New(app)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
