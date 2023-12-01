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

type CategoryViews struct {
	CategoryName string `json:"categoryName"`
	Views        int    `json:"views"`
}

type GetAdStatsResponse struct {
	AdvertisingId         string          `json:"advertisingId"`
	TotalViews            int             `json:"totalViews"`
	TotalViewsPerCategory []CategoryViews `json:"totalViewsPerCategory"`
	DailyViews            []DailyView     `json:"dailyViews"`
}

func (h GetStatsHandler) handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	advertisingId := req.PathParameters["advertisingId"]
	if advertisingId == "" {
		log.Fatalf("AdvertisingId is empty")
	}

	var totalViews = 0
	views := h.viewsRepo.FindViewsByAdId(advertisingId)

	dayToCategories := make(map[string]map[string]int)
	totalViewsPerCategoryMap := make(map[string]int)

	for _, view := range views {
		dateString := utils.UnixTimestampToDateString(view.Timestamp)
		var viewsPerCategory = dayToCategories[dateString]

		if viewsPerCategory == nil {
			viewsPerCategory = make(map[string]int)
		}

		for categoryUuid, viewsAmount := range view.Views {
			category := utils.CategoryUuidToString(categoryUuid)
			amountOfViews := viewsPerCategory[category]
			viewsPerCategory[category] = amountOfViews + viewsAmount //TODO: Rename it

			categoryViews := totalViewsPerCategoryMap[category]
			totalViewsPerCategoryMap[category] = categoryViews + viewsAmount

			totalViews = totalViews + amountOfViews
		}

		dayToCategories[dateString] = viewsPerCategory
	}

	// Converts map category -> total views to a json readable object
	flatCategoryViews := func(categoriesMap map[string]int) []CategoryViews {
		flatArray := make([]CategoryViews, 0)
		for categoryName, categoryViews := range categoriesMap {
			flatArray = append(flatArray, CategoryViews{CategoryName: categoryName, Views: categoryViews})
		}
		return flatArray
	}(totalViewsPerCategoryMap)

	dailyViews := make([]DailyView, 0)
	for day, categoryMap := range dayToCategories {
		dailyViews = append(dailyViews, DailyView{Date: day, ViewsPerCategory: categoryMap})
	}

	return transport.SendOk(GetAdStatsResponse{
		TotalViews:            totalViews,
		AdvertisingId:         advertisingId,
		DailyViews:            dailyViews,
		TotalViewsPerCategory: flatCategoryViews,
	}), nil
}

func main() {
	db := database.InitDynamo()
	viewsRepo := database.NewViewsRepo(db)
	getAdStatsHandler := GetStatsHandler{viewsRepo: viewsRepo}
	lambda.Start(getAdStatsHandler.handleRequest)
}
