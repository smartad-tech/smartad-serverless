package database

import "github.com/smartad-tech/smartad-serverless/internal/types"

type ViewEntity struct {
	AdvertisingId string                     `dynamodbav:"advertising_id"`
	Timestamp     string                     `dynamodbav:"timestamp"`
	Views         map[types.CategoryUuid]int `dynamodbav:"views"`
}
