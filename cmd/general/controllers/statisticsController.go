package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smartad-tech/smartad-serverless/internal/database"
)

type StatisticsController struct {
	viewsRepository database.IViewsRepository
}

type StatsCategoryView struct {
	CategoryName string `json:"categoryName"`
	Views        int    `json:"views"`
}

type DailyView struct {
	Date  string              `json:"date"`
	Views []StatsCategoryView `json:"views"`
}

func (c StatisticsController) GetDailyViews(ctx *fiber.Ctx) error {
	// from := ctx.Query("from")
	// to := ctx.Query("to")
	// adId := ctx.Params("adId")
	// if from == "" || to == "" {
	// 	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": "`from` or `to` were not provided.",
	// 	})
	// }

	// fromDate, toDate, err := func(from string, to string) (utils.SmartDate, utils.SmartDate, error) {
	// 	fromDate, fromError := utils.NewSmartDateFromString(from)
	// 	toDate, toError := utils.NewSmartDateFromString(to)
	// 	err := errors.Join(fromError, toError)
	// 	if err != nil {
	// 		return utils.SmartDate{}, utils.SmartDate{}, err
	// 	}
	// 	return fromDate, toDate, nil
	// }(from, to)
	// if err != nil {
	// 	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": "Wrong date formats provided",
	// 	})
	// }

	// views, err := c.viewsRepository.FindViewsByAdIdInRange(adId, fromDate.Unix(), toDate.Unix())
	// if err != nil {
	// 	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": "Error during fetching database.",
	// 	})
	// }

	// dateToViewsMap := make(map[types.Date]map[types.CategoryName]int)
	// for _, viewEntity := range views {
	// 	dayDate, err := utils.NewSmartDateFromUnix(viewEntity.Timestamp)
	// 	if err != nil {
	// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 			"error": "Could not aggregate views data",
	// 		})
	// 	}

	// 	categoryToViewsMap := dateToViewsMap[dayDate.ToString()]
	// 	if categoryToViewsMap == nil {
	// 		categoryToViewsMap = make(map[types.CategoryName]int)
	// 	}

	// 	for categoryUuid, amountOfViews := range viewEntity.Views {
	// 		categoryName := utils.CategoryUuidToString(categoryUuid)
	// 		categoryViews := categoryToViewsMap[categoryName]
	// 		categoryToViewsMap[categoryName] = categoryViews + amountOfViews
	// 	}
	// 	dateToViewsMap[dayDate.ToString()] = categoryToViewsMap
	// }

	// dailyViews := make([]DailyView, 0)
	// for dateString, categoryToViewsMap := range dateToViewsMap {
	// 	statsCategoryViews := make([]StatsCategoryView, 0)

	// 	for key, value := range categoryToViewsMap {
	// 		statsCategoryViews = append(statsCategoryViews, StatsCategoryView{
	// 			CategoryName: key,
	// 			Views:        value,
	// 		})
	// 	}

	// 	dailyViews = append(dailyViews, DailyView{
	// 		Date:  dateString,
	// 		Views: statsCategoryViews,
	// 	})
	// }

	// return ctx.JSON(dailyViews)
	return ctx.SendStatus(200)
}

func NewStatisticsController(viewsRepository database.IViewsRepository) StatisticsController {
	return StatisticsController{
		viewsRepository: viewsRepository,
	}
}
