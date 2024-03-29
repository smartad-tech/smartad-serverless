package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smartad-tech/smartad-serverless/cmd/general/factory"
)

func PublicRoutes(a *fiber.App, controllers factory.Controllers) {
	// Create routes group.
	baseRoute := a.Group("/api/v1")

	// Routes for GET method:
	baseRoute.Get("smartads/:adId/views/daily", controllers.DashboardController.GetDailyViews)
	baseRoute.Post("smartads/:adId/views", controllers.DeviceController.PostAdViews)
	baseRoute.Get("smartads/:adUuid/views/length", controllers.DashboardController.GetAvgViewLength)
}
