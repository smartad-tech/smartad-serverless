package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/smartad-tech/smartad-serverless/internal/database"
	"github.com/smartad-tech/smartad-serverless/internal/transport"
	"github.com/smartad-tech/smartad-serverless/internal/utils"
)

type GetStatsHandler struct {
	viewsRepo *database.ViewsRepository
}

type DailyView struct {
	Date             string         `json:"date"`
	ViewsPerCategory map[string]int `json:"viewsPerCategory"`
}

type GetAdStatsResponse struct {
	AdvertisingId string      `json:"advertisingId"`
	TotalViews    int         `json:"totalViews"`
	DailyViews    []DailyView `json:"dailyViews"`
}

func (h GetStatsHandler) handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	advertisingId := req.PathParameters["advertisingId"]
	if advertisingId == "" {
		log.Fatalf("AdvertisingId is empty")
	}

	var totalViews = 0
	views := h.viewsRepo.FindViewsByAdId(advertisingId)

	dayToCategories := make(map[string]map[string]int)

	for _, view := range views {
		dateString := utils.UnixTimestampToDateString(view.Timestamp)
		var viewsPerCategory = dayToCategories[dateString]

		if viewsPerCategory == nil {
			viewsPerCategory = make(map[string]int)
		}

		for category, viewsAmount := range view.Views {
			amountOfViews := viewsPerCategory[category]
			viewsPerCategory[category] = amountOfViews + viewsAmount //TODO: Rename it
			totalViews = totalViews + amountOfViews
		}

		dayToCategories[dateString] = viewsPerCategory
	}

	dailyViews := make([]DailyView, 0)

	for day, categoryMap := range dayToCategories {
		dailyViews = append(dailyViews, DailyView{Date: day, ViewsPerCategory: categoryMap})
	}

	return transport.SendOk(GetAdStatsResponse{
		TotalViews:    totalViews,
		AdvertisingId: advertisingId,
		DailyViews:    dailyViews,
	}), nil
}

func main() {
	db := database.InitDynamo()
	viewsRepo := database.NewViewsRepo(db)
	getAdStatsHandler := GetStatsHandler{viewsRepo: viewsRepo}
	lambda.Start(getAdStatsHandler.handleRequest)
}
