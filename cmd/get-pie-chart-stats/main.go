package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/smartad-tech/smartad-serverless/cmd/get-pie-chart-stats/handler"
	"github.com/smartad-tech/smartad-serverless/cmd/get-pie-chart-stats/service"
	"github.com/smartad-tech/smartad-serverless/internal/database"
)

func main() {
	db := database.InitDynamo()
	viewsRepo := database.NewViewsRepo(db)
	getAdService := service.NewService(*viewsRepo)
	getAdStatsHandler := handler.NewHandler(*getAdService)
	lambda.Start(getAdStatsHandler.Handle)
}
