package db

import (
  "os"
  "github.com/jinzhu/gorm"
  "github.com/joho/godotenv"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var Connection *gorm.DB

func Connect() {
  godotenv.Load()
  connectionString := os.Getenv("CONNECTION_STRING")
  dbConnection, err := gorm.Open("postgres", connectionString)

  if err != nil {
    panic(err)
  }

  Connection = dbConnection
}
