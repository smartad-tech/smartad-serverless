package service

import (
	"github.com/smartad-tech/smartad-serverless/internal/database"
	"github.com/smartad-tech/smartad-serverless/internal/transport"
	"github.com/smartad-tech/smartad-serverless/internal/utils"
)

type CategoryName = string

type IGetAdService interface {
	GetStatsByAdId(advertisingId string) ([]transport.CategoryViews, error)
}

type GetAdService struct {
	viewsRepo database.IViewsRepository
}

func (s GetAdService) GetStatsByAdId(advertisingId string) ([]transport.CategoryViews, error) {
	var categoryViews []transport.CategoryViews
	views, err := s.viewsRepo.FindViewsByAdId(advertisingId)
	if err != nil {
		return categoryViews, err
	}

	func(viewEntities []database.ViewEntity, categoryViews *[]transport.CategoryViews) {
		categoryViewsMap := make(map[CategoryName]int)
		for _, viewEntity := range views {
			for categoryUuid, numberOfViews := range viewEntity.Views {
				categoryName := utils.CategoryUuidToString(categoryUuid)
				categoryNumberOfViews := categoryViewsMap[categoryName]
				categoryViewsMap[categoryName] = categoryNumberOfViews + numberOfViews
			}
		}

		for categoryName, amountOfViews := range categoryViewsMap {
			*categoryViews = append(*categoryViews, transport.CategoryViews{
				CategoryName: categoryName,
				Views:        amountOfViews,
			})
		}
	}(views, &categoryViews)

	return categoryViews, nil
}

func NewService(viewsRepo database.IViewsRepository) *GetAdService {
	return &GetAdService{viewsRepo: viewsRepo}
}
