package service_test

import (
	"github.com/smartad-tech/smartad-serverless/internal/database"
	"testing"

	"github.com/smartad-tech/smartad-serverless/cmd/get-pie-chart-stats/service"
)

const AdultMenCategoryUuid = "c44453a1-8184-4905-a81e-04372f9b76e6"
const AdultWomenCategoryUuid = "7186a646-17ec-4afb-a18f-104c34830eac"

type ViewDatabaseMock struct {
}

func (r ViewDatabaseMock) FindViewsByAdId(_ string) ([]database.ViewEntity, error) {
	views := []database.ViewEntity{
		{Views: map[string]int{AdultMenCategoryUuid: 3}, Timestamp: "any", AdvertisingId: "123"},
		{Views: map[string]int{AdultWomenCategoryUuid: 10}, Timestamp: "any", AdvertisingId: "123"},
		{Views: map[string]int{AdultMenCategoryUuid: 15, AdultWomenCategoryUuid: 30}, Timestamp: "any", AdvertisingId: "123"},
	}
	return views, nil
}

func TestGetStatsByAdId(t *testing.T) {
	getStatsService := service.NewService(ViewDatabaseMock{})
	categoryViews, err := getStatsService.GetStatsByAdId("123")
	if err != nil {
		t.Fail()
	}
	size := len(categoryViews)
	if size != 2 {
		t.Fail()
	}

	for _, categoryView := range categoryViews {
		switch categoryView.CategoryName {
		case "Adult Men":
			if categoryView.Views != 18 {
				t.Fail()
			}
			break
		case "Adult Women":
			if categoryView.Views != 40 {
				t.Fail()
			}
			break
		default:
			t.Fail()
			break
		}
	}
}
