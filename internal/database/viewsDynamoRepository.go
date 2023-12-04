package database

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/jsii-runtime-go"
)

type IViewsRepository interface {
	FindViewsByAdId(advertisingId string) ([]ViewEntity, error)
}

type ViewsRepository struct {
	dynamoClient *dynamodb.Client
}

const tableName = "smartad-views-table"

func (r ViewsRepository) FindViewsByAdId(advertisingId string) ([]ViewEntity, error) {
	entities := make([]ViewEntity, 0)
	keyCondition := expression.Key("advertising_id").Equal(expression.Value(advertisingId))

	expr, expressionError := expression.NewBuilder().WithKeyCondition(keyCondition).Build()
	if expressionError != nil {
		log.Printf("Expression builder returned error. Error: %s", expressionError.Error())
		return entities, expressionError
	}

	response, queryError := r.dynamoClient.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 jsii.String(tableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})
	if queryError != nil {
		log.Printf("FindViewsByAdId returned an error. Error: %s", queryError.Error())
		return entities, queryError
	}

	for _, item := range response.Items {
		view := ViewEntity{}
		decodingError := attributevalue.UnmarshalMap(item, &view)
		if decodingError != nil {
			log.Fatalf("Got error unmarshalling item: %s", decodingError)
		}

		entities = append(entities, view)
	}
	return entities, nil
}

func NewViewsRepo(dynamoClient *dynamodb.Client) *ViewsRepository {
	return &ViewsRepository{
		dynamoClient: dynamoClient,
	}
}
