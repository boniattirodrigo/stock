package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var Connection *gorm.DB

func Connect() {
	godotenv.Load()
	connectionString := os.Getenv("CONNECTION_STRING")
	dbConnection, err := gorm.Open("postgres", connectionString)

	if os.Getenv("ENVIRONMENT") == "development" {
		dbConnection.LogMode(true)
	}

	if err != nil {
		panic(err)
	}

	Connection = dbConnection
}
