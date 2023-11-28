package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/smartad-tech/smartad-serverless/internal/transport"
	"log"
)

type DemoLoginResponse struct {
	AdvertisingId string `json:"advertisingId"`
	UserId        string `json:"userId"`
}

func handle(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	pin := request.QueryStringParameters["pin"]
	if pin == "" {
		log.Print("Didn't receive PIN. Rejecting...")
		return transport.SendBadRequest("Invalid/Empty PIN."), nil
	}

	if pin == "1337" {
		demoLoginResponse := DemoLoginResponse{
			AdvertisingId: "123",
			UserId:        "66709f40-b39e-41dc-b118-4904681c1572", // Random UUID as a replacement
		}
		return transport.SendOk(demoLoginResponse), nil
	}

	return transport.SendNotAuthorized(), nil
}

func main() {
	lambda.Start(handle)
}
