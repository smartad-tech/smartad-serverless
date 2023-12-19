package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/smartad-tech/smartad-serverless/internal/types"
)

type DeviceController struct {
}

func NewDeviceController() DeviceController {
	return DeviceController{}
}

type PostViewsRequest struct {
	CategoryViewsMap map[types.CategoryUuid]int `json:"categoryViews"`
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

	return ctx.SendStatus(200)
}
