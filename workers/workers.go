package workers

import (
	"fmt"
	"time"
	"strconv"
  "strings"
  "math/rand"
  "github.com/boniattirodrigo/stock/ws"
  "github.com/boniattirodrigo/stock/models"
  "github.com/boniattirodrigo/stock/db"
  "github.com/gocolly/colly"
)

func updateRandomStockPrince(ticker string) {
  var stock models.Stock
  db.Connection.Where("ticker = ?", ticker).First(&stock)
  db.Connection.Model(&stock).Update("price", rand.Float64() * 10)
  ws.StockPublisher()
}

func updateStockPrice(ticker string) {
  timer := time.NewTimer(5 * time.Second)
  <-timer.C

  c := colly.NewCollector()

  c.OnHTML(".special strong", func(e *colly.HTMLElement) {
    var stock models.Stock
    price, err := strconv.ParseFloat(strings.ReplaceAll(e.Text, ",", "."), 64)

    if err == nil {
      db.Connection.Where("ticker = ?", ticker).First(&stock)
      db.Connection.Model(&stock).Update("price", price)
      ws.StockPublisher()
    }
  })

  c.Visit(fmt.Sprint("https://statusinvest.com.br/acoes/", ticker))
}

func Start() {
  var stocks []models.Stock
  var tickers []string
  db.Connection.Find(&stocks).Pluck("ticker", &tickers)

  // FOR DEVELOPMENT
  // for _, ticker := range tickers {
  //   timer := time.NewTimer(5 * time.Second)
  //   <-timer.C

  //   go updateRandomStockPrince(ticker)
  // }

  // FOR PRODUCTION
  timeTicker := time.NewTicker(8 * time.Minute)

  for {
    select {
    case <-timeTicker.C:
      for _, ticker := range tickers {
        timer := time.NewTimer(5 * time.Second)
        <-timer.C

        go updateStockPrice(ticker)
      }
    }
  }
}
