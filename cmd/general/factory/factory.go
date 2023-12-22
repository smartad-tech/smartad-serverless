package factory

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/smartad-tech/smartad-serverless/cmd/general/controllers"
	"github.com/smartad-tech/smartad-serverless/internal/database"
)

type Controllers struct {
	StatisticsController controllers.StatisticsController
	DeviceController     controllers.DeviceController
}

type Repositories struct {
	ViewsRepository database.IViewsRepository
}

func InitControllers(repositories Repositories) Controllers {
	statsController := controllers.NewStatisticsController(repositories.ViewsRepository)
	deviceController := controllers.NewDeviceController(repositories.ViewsRepository)
	return Controllers{
		StatisticsController: statsController,
		DeviceController:     deviceController,
	}
}

func InitRepositories(dynamoClient *dynamodb.Client) Repositories {
	viewsRepository := database.NewViewsRepo(dynamoClient)
	return Repositories{
		ViewsRepository: viewsRepository,
	}
}
