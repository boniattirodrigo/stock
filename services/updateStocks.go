package services

import (
	"fmt"
	"strconv"
	"strings"
  "github.com/boniattirodrigo/stock/models"
  "github.com/gocolly/colly"
  "github.com/jinzhu/gorm"
)

func UpdateStocks(stockName string, dbConnection *gorm.DB) {
  c := colly.NewCollector()

	c.OnHTML(".special strong", func(e *colly.HTMLElement) {
    var stock models.Stock
    value, err := strconv.ParseFloat(strings.ReplaceAll(e.Text, ",", "."), 64)

    if err == nil {
      dbConnection.Where("ticket = ?", stockName).First(&stock)
      dbConnection.Model(&stock).Update("price", value)
      fmt.Println(stockName, e.Text)
    }
	})

  c.Visit(fmt.Sprint("https://statusinvest.com.br/acoes/", stockName))
}
