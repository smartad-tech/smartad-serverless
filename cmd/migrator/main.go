package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/smartad-tech/smartad-serverless/internal/database"
	"github.com/smartad-tech/smartad-serverless/internal/transport"
	"gorm.io/gorm"
)

var db *gorm.DB

func migrateLambdaHandler() (events.APIGatewayProxyResponse, error) {
	log.Print("Starting migration lambda...")
	err := db.AutoMigrate(database.AdEntity{}, database.DeviceEntity{}, database.UserEntity{})
	if err != nil {
		log.Printf("Error during neon migration. Error: %s", err.Error())
		return transport.SendServerError(), err
	}
	log.Print("The migration lambda was successfully executed!")
	return transport.SendOk("OK"), nil
}

func main() {
	sqlConnection, err := database.InitSql()
	if err != nil {
		log.Print("Failed to connect to database. Exiting migrator lambda", err.Error())
		os.Exit(1)
	}
	db = sqlConnection
	lambda.Start(migrateLambdaHandler)
}
