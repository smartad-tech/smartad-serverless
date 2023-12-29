package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitSql() (*gorm.DB, error) {
	password := os.Getenv("NEON_DB_PASSWORD")
	username := os.Getenv("NEON_DB_USERNAME")
	host := os.Getenv("NEON_DB_HOST")
	database := os.Getenv("NEON_DB_NAME")
	dbs := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=require", username, password, host, database)

	db, err := gorm.Open(postgres.Open(dbs), &gorm.Config{})
	if err != nil {
		log.Print("Error during initializing connection to postgres db.", err.Error())
		return db, err
	}
	return db, nil
}
