package controllers

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/smartad-tech/smartad-serverless/internal/database"
	"github.com/smartad-tech/smartad-serverless/internal/types"
	"github.com/smartad-tech/smartad-serverless/internal/utils"
)

type DeviceController struct {
	repository database.IViewsRepository
}

func NewDeviceController(repo database.IViewsRepository) DeviceController {
	return DeviceController{
		repository: repo,
	}
}

type CategoryView struct {
	CategoryUuid types.CategoryUuid `json:"categoryUuid"`
	Amount       int                `json:"amount"`
}

type PostViewsRequest struct {
	CategoryViews []CategoryView `json:"categoryViews"`
}

func (c DeviceController) PostViews(ctx *fiber.Ctx) error {
	adId := ctx.Params("adId")
	if adId == "" {
		return ctx.SendStatus(400)
	}
	postViewsRequest := PostViewsRequest{}
	err := json.Unmarshal(ctx.Body(), &postViewsRequest)
	if err != nil {
		return ctx.SendStatus(400)
	}

	now := utils.NewSmartDate()

	viewsMap := make(map[types.CategoryUuid]int)

	for _, categoryView := range postViewsRequest.CategoryViews {
		// if category uuid is empty. Later add that to validation
		if categoryView.CategoryUuid == "" {
			return ctx.SendStatus(fiber.ErrBadRequest.Code)
		}
		amountOfViews := viewsMap[categoryView.CategoryUuid]
		viewsMap[categoryView.CategoryUuid] = amountOfViews + categoryView.Amount
	}

	err = c.repository.SaveViews(adId, viewsMap, now.Unix())
	if err != nil {
		log.Printf("Error during saving views. Error: %s", err.Error())
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
