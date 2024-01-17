package factory

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dashboardcontr "github.com/smartad-tech/smartad-serverless/cmd/general/controllers/dashboard-controller"
	devicecontr "github.com/smartad-tech/smartad-serverless/cmd/general/controllers/device-controller"
	"github.com/smartad-tech/smartad-serverless/internal/database"
)

type Controllers struct {
	DashboardController dashboardcontr.DashboardController
	DeviceController    devicecontr.DeviceController
}

type Repositories struct {
	ViewsRepository database.IViewsRepository
}

func InitControllers(repositories Repositories) Controllers {
	statsController := dashboardcontr.NewStatisticsController(repositories.ViewsRepository)
	deviceController := devicecontr.NewDeviceController(repositories.ViewsRepository)
	return Controllers{
		DashboardController: statsController,
		DeviceController:    deviceController,
	}
}

func InitRepositories(dynamoClient *dynamodb.Client) Repositories {
	viewsRepository := database.NewViewsRepo(dynamoClient)
	return Repositories{
		ViewsRepository: viewsRepository,
	}
}
