package database

import "github.com/smartad-tech/smartad-serverless/internal/types"

type ViewEntity struct {
	AdvertisingUuid string             `dynamodbav:"advertising_uuid"`
	CategoryUuid    types.CategoryUuid `dynamodbav:"category_uuid"`
	Timestamp       string             `dynamodbav:"timestamp"`
	ViewLength      float32            `dynamodbav:"view_length"`
	DeviceUuid      string             `dynamodbav:"device_uuid"`
}
