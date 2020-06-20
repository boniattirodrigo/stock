package resolvers

import (
	"time"
	"github.com/boniattirodrigo/stock/db"
  "github.com/boniattirodrigo/stock/models"
)

type Stock struct {
  Id         uint
  Ticker     string
  Price      float64
  CreatedAt  time.Time
  UpdatedAt  time.Time
}

func FetchStocks(tickers []string) ([]Stock) {
  var stocks []models.Stock
  db.Connection.Where("ticker IN (?)", tickers).Find(&stocks)
  var stockData = []Stock{}

  for _, stock := range stocks {
    newStock := Stock{ Id: stock.ID, Ticker: stock.Ticker, Price: stock.Price, CreatedAt: stock.CreatedAt, UpdatedAt: stock.UpdatedAt }
    stockData = append(stockData, newStock)
  }

  return stockData
}
