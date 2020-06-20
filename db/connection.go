package db

import (
  "log"
  "os"
  "github.com/joho/godotenv"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var Connection *gorm.DB

func Connect() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  connectionString := os.Getenv("CONNECTION_STRING")
  dbConnection, err := gorm.Open("postgres", connectionString)

  if err != nil {
    panic(err)
  }

  Connection = dbConnection
}
