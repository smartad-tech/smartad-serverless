package devicecontr

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/smartad-tech/smartad-serverless/internal/types"
	"github.com/smartad-tech/smartad-serverless/internal/utils"
)

type PostAdViewsBody struct {
	Views      []ViewBody `json:"views"`
	DeviceUuid string     `json:"deviceId"`
}

type ViewBody struct {
	ViewLength   float32            `json:"viewLength"`
	CategoryUuid types.CategoryUuid `json:"categoryUuid"`
}

func (c DeviceController) PostAdViews(ctx *fiber.Ctx) error {
	adId := ctx.Params("adId")
	if adId == "" {
		return ctx.SendStatus(400)
	}
	postAdViewsBody := PostAdViewsBody{}
	err := json.Unmarshal(ctx.Body(), &postAdViewsBody)
	if err != nil {
		return ctx.SendStatus(400)
	}

	now := utils.NewSmartDate()

	var isFailed = false // Used to track if one of views were not saved
	for _, viewBody := range postAdViewsBody.Views {
		err = c.repository.SaveView(adId, viewBody.CategoryUuid, now.Unix(), postAdViewsBody.DeviceUuid, viewBody.ViewLength)
		if err != nil {
			log.Printf("Error during saving views. Error: %s", err.Error())
			log.Printf("Failed to save view: [%+v] for device: [%s]", viewBody, postAdViewsBody.DeviceUuid)
			isFailed = true
		}
	}

	if isFailed {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
