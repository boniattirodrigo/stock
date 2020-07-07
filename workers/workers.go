package workers

import (
	"fmt"
	"github.com/boniattirodrigo/stock/db"
	"github.com/boniattirodrigo/stock/models"
	"github.com/boniattirodrigo/stock/ws"
	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func updateRandomStockPrice(ticker string) {
	var stock models.Stock
	db.Connection.Model(&stock).Where("ticker = ?", ticker).Update("price", rand.Float64()*10)
	ws.StockPublisher()
}

func updateStockPrice(ticker string) {
	c := colly.NewCollector()

	c.OnHTML(".special strong", func(e *colly.HTMLElement) {
		var stock models.Stock
		price, err := strconv.ParseFloat(strings.ReplaceAll(e.Text, ",", "."), 64)

		if err == nil {
			db.Connection.Model(&stock).Where("ticker = ?", ticker).Update("price", price)
			ws.StockPublisher()
		}
	})

	c.Visit(fmt.Sprint("https://statusinvest.com.br/acoes/", ticker))
}

func Start() {
	var stocks []models.Stock
	var tickers []string
	godotenv.Load()
	db.Connection.Find(&stocks).Pluck("ticker", &tickers)

	if os.Getenv("ENVIRONMENT") == "development" {
		for _, ticker := range tickers {
			timer := time.NewTimer(5 * time.Second)
			<-timer.C

			go updateRandomStockPrice(ticker)
		}
	} else {
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
}
