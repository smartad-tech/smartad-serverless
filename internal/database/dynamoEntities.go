package database

type ViewEntity struct {
	AdvertisingId string         `dynamodbav:"advertising_id"`
	Timestamp     string         `dynamodbav:"timestamp"`
	Views         map[string]int `dynamodbav:"views"`
}
