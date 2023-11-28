package database

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/jsii-runtime-go"
)

type ViewsRepository struct {
	dynamoClient *dynamodb.Client
}

var tableName = "smartad-views-table"

func (r ViewsRepository) FindViewsByAdId(advertisingId string) []ViewEntity {
	keyCondition := expression.Key("advertising_id").Equal(expression.Value(advertisingId))

	expr, expressionError := expression.NewBuilder().WithKeyCondition(keyCondition).Build()
	if expressionError != nil {
		log.Fatalf("Expression builder returned error. Error: %s", expressionError.Error())
	}

	response, queryError := r.dynamoClient.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 jsii.String(tableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})
	if queryError != nil {
		log.Fatalf("FindViewsByAdId returned an error. Error: %s", queryError.Error())
	}

	var views []ViewEntity

	for _, item := range response.Items {
		view := ViewEntity{}
		decodingError := attributevalue.UnmarshalMap(item, &view)
		if decodingError != nil {
			log.Fatalf("Got error unmarshalling item: %s", decodingError)
		}

		views = append(views, view)
	}
	return views
}

func NewViewsRepo(dynamoClient *dynamodb.Client) *ViewsRepository {
	return &ViewsRepository{
		dynamoClient: dynamoClient,
	}
}
