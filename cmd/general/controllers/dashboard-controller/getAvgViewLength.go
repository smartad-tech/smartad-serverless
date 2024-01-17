package dashboardcontr

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type GetAvgViewLengthResponse struct {
	AdvertisingUuid   string  `json:"advertisingUuid"`
	AverageViewLength float32 `json:"averageViewLength"`
}

func (c DashboardController) GetAvgViewLength(ctx *fiber.Ctx) error {
	adUuid := ctx.Params("adUuid")
	if adUuid == "" {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	views, err := c.viewsRepository.FindAllViews(adUuid)
	if err != nil {
		log.Printf("Got an error during searching for all views. Error: %s", err.Error())
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	var sum float32 = 0.0
	for _, view := range views {
		sum += view.ViewLength
	}
	var avgViewLength = sum / float32(len(views))

	return ctx.Status(fiber.StatusOK).JSON(GetAvgViewLengthResponse{AdvertisingUuid: adUuid, AverageViewLength: avgViewLength})
}
