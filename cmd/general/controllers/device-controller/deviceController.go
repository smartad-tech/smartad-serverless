package devicecontr

import (
	"github.com/smartad-tech/smartad-serverless/internal/database"
)

type DeviceController struct {
	repository database.IViewsRepository
}

func NewDeviceController(repo database.IViewsRepository) DeviceController {
	return DeviceController{
		repository: repo,
	}
}
