package database

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var defaultRegion = "eu-central-1"

func InitDynamo() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(defaultRegion))
	if err != nil {
		log.Panicf("Error during initializing dynamodb config. Error: %s", err.Error())
	}
	return dynamodb.NewFromConfig(cfg)
}
