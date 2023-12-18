package controllers

import "github.com/gofiber/fiber/v2"

type DeviceController struct {
}

func NewDeviceController() DeviceController {
	return DeviceController{}
}

func (c DeviceController) PostViews(ctx *fiber.Ctx) error {
	return ctx.SendStatus(200)
}
