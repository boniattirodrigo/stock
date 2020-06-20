package db

import (
  "github.com/boniattirodrigo/stock/models"
  "github.com/jinzhu/gorm"
)

func RunMigrations(stocks [3]string, dbConnection *gorm.DB) {
  dbConnection.AutoMigrate(&models.Stock{})

  for _, stockName := range stocks {
    var stock models.Stock
    dbConnection.FirstOrCreate(&stock, models.Stock{Ticket: stockName})
  }
}
