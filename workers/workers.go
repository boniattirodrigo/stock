package workers

import (
	"fmt"
	"github.com/boniattirodrigo/stock/db"
	"github.com/boniattirodrigo/stock/models"
	"github.com/boniattirodrigo/stock/ws"
	"github.com/gocolly/colly"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func updateRandomStockPrice(ticker string) {
	var stock models.Stock
	db.Connection.Model(&stock).Where("ticker = ?", ticker).Update("price", rand.Float64()*10)
	ws.StockPublisher()
}

func updateStockPrice(url string, selector string, ticker string) {
	c := colly.NewCollector()

	c.OnHTML(selector, func(e *colly.HTMLElement) {
		var stock models.Stock
		price, err := strconv.ParseFloat(strings.ReplaceAll(e.Text, ",", "."), 64)

		if err == nil {
			db.Connection.Model(&stock).Where("ticker = ?", ticker).Update("price", price)
			ws.StockPublisher()
		}
	})

	c.Visit(fmt.Sprint(url, ticker))
}

func crawlPages(url string, selector string, tickers []string, interval int) {
	readyToCrawlChannel := make(chan bool)
	totalCrawled := 0

	go func() {
		readyToCrawlChannel <- true
	}()

	for {
		select {
		case <-readyToCrawlChannel:
			for _, ticker := range tickers {
				timer := time.NewTimer(time.Duration(interval) * time.Second)
				<-timer.C

				go updateStockPrice(url, selector, ticker)

				totalCrawled += 1

				if len(tickers) == totalCrawled {
					go func() {
						totalCrawled = 0
						readyToCrawlChannel <- true
					}()
				}
			}
		}
	}
}

func Start() {
	var stocks []models.Stock
	var tickersAsc []string
	var tickersDesc []string
	db.Connection.Order("ticker asc").Find(&stocks).Pluck("ticker", &tickersAsc)
	db.Connection.Order("ticker desc").Find(&stocks).Pluck("ticker", &tickersDesc)

	go crawlPages("https://statusinvest.com.br/acoes/", ".special strong", tickersAsc, 5)
	go crawlPages("https://www.infomoney.com.br/", ".line-info .value p", tickersDesc, 1)
}
