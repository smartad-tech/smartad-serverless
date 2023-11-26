package transport

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

func SendOk(body any) events.APIGatewayProxyResponse {
	jsonBytes, _ := json.Marshal(body)
	return events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           map[string]string{"content-type": "application/json", "Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "*"},
		MultiValueHeaders: nil,
		Body:              string(jsonBytes),
		IsBase64Encoded:   false,
	}
}

func SendBadRequest(errorMessage string) events.APIGatewayProxyResponse {
	jsonBytes, _ := json.Marshal(errorMessage)
	return events.APIGatewayProxyResponse{
		StatusCode:        400,
		Headers:           map[string]string{"content-type": "application/json", "Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "*"},
		MultiValueHeaders: nil,
		Body:              string(jsonBytes),
		IsBase64Encoded:   false,
	}
}

func SendNotAuthorized() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode:        401,
		Headers:           map[string]string{"content-type": "application/json", "Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "*"},
		MultiValueHeaders: nil,
		Body:              "",
		IsBase64Encoded:   false,
	}
}
