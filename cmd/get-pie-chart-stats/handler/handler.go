package handler

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/smartad-tech/smartad-serverless/cmd/get-pie-chart-stats/service"
	"github.com/smartad-tech/smartad-serverless/internal/transport"
	"log"
)

type GetPieChartStatsHandler struct {
	service service.IGetAdService
}

func (h GetPieChartStatsHandler) Handle(req events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	advertisingId := req.PathParameters["advertisingId"]
	if advertisingId == "" {
		log.Print("Received an empty advertisingId. BAD_REQUEST")
		return transport.SendBadRequest("No advertisingId provided"), nil
	}
	categoryViews, err := h.service.GetStatsByAdId(advertisingId)
	if err != nil {
		log.Printf("Received an error after getting stats. INTERNAL_SERVER_ERROR")
		return transport.SendServerError(), nil
	}

	return transport.SendOk(transport.GetAdStatsResponse{
		AdvertisingId:         advertisingId,
		TotalViewsPerCategory: categoryViews,
	}), nil
}

func NewHandler(getAdService service.GetAdService) *GetPieChartStatsHandler {
	return &GetPieChartStatsHandler{
		service: getAdService,
	}
}
