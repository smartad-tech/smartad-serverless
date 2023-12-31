package database

import (
	"context"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/jsii-runtime-go"
	"github.com/smartad-tech/smartad-serverless/internal/types"
)

type UnixTime = int64

const tableName = "smartad-views-table"

type IViewsRepository interface {
	FindViewsByAdId(advertisingId string) ([]ViewEntity, error)
	FindViewsByAdIdInRange(advertisingId string, from UnixTime, to UnixTime) ([]ViewEntity, error)
	SaveViews(advertisingId string, categoryToViewsMap map[types.CategoryUuid]int, timestamp UnixTime) error
}

type ViewsRepository struct {
	dynamoClient *dynamodb.Client
}

func (r ViewsRepository) FindViewsByAdIdInRange(advertisingId string, from UnixTime, to UnixTime) ([]ViewEntity, error) {
	var entities []ViewEntity
	fromString := strconv.FormatInt(from, 10)
	toString := strconv.FormatInt(to, 10)
	log.Printf("From [%s] to [%s]", fromString, toString)
	keyCondition := expression.Key("advertising_id").Equal(expression.Value(advertisingId)).And(expression.Key("timestamp").Between(expression.Value(fromString), expression.Value(toString)))

	expr, expressionErr := expression.NewBuilder().WithKeyCondition(keyCondition).Build()
	if expressionErr != nil {
		log.Printf("Error during building query views in range. Error: %s", expressionErr.Error())
		return entities, expressionErr
	}

	response, queryError := r.dynamoClient.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 jsii.String(tableName),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})
	if queryError != nil {
		log.Printf("Error during quering views in range. Error: %s", queryError.Error())
		return entities, queryError
	}

	for _, item := range response.Items {
		view := ViewEntity{}
		decodingError := attributevalue.UnmarshalMap(item, &view)
		if decodingError != nil {
			log.Printf("Got error unmarshalling item: %s", decodingError)
			return entities, decodingError
		}

		entities = append(entities, view)
	}
	log.Printf("Found %d view entities in given range.", len(entities))
	return entities, nil
}

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

func (r ViewsRepository) SaveViews(advertisingId string, categoryToViewsMap map[types.CategoryUuid]int, timestamp UnixTime) error {
	viewEntity := ViewEntity{
		AdvertisingId: advertisingId,
		Timestamp:     strconv.FormatInt(timestamp, 10),
		Views:         categoryToViewsMap,
	}

	item, err := attributevalue.MarshalMap(viewEntity)
	if err != nil {
		log.Printf("Error during marshalling view entity for saving views. Error: %s", err.Error())
		return err
	}

	_, err = r.dynamoClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: jsii.String(tableName),
		Item:      item,
	})
	if err != nil {
		log.Printf("Error during putItem. Error: %s", err.Error())
		return err
	}

	return nil
}

func NewViewsRepo(dynamoClient *dynamodb.Client) *ViewsRepository {
	return &ViewsRepository{
		dynamoClient: dynamoClient,
	}
}
