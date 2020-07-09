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

func updateStockPrice(url string, selector string, ticker string, timeout time.Duration) {
	c := colly.NewCollector()

	c.SetRequestTimeout(timeout)

	c.OnHTML(selector, func(e *colly.HTMLElement) {
		var stock models.Stock
		price, err := strconv.ParseFloat(strings.ReplaceAll(e.Text, ",", "."), 64)

		if err == nil {
			db.Connection.Where("ticker = ?", ticker).First(&stock)

			if stock.Price != price {
				stock.Price = price
				db.Connection.Save(&stock)
				ws.StockPublisher()
			}
		}
	})

	c.Visit(fmt.Sprint(url, ticker))
}

func crawlPages(url string, selector string, tickers []string, interval int) {
	readyToCrawlChannel := make(chan bool)
	readyToCrawl := func() { readyToCrawlChannel <- true }
	go readyToCrawl()

	for {
		select {
		case <-readyToCrawlChannel:
			for index, ticker := range tickers {
				intervalTime := time.Duration(interval) * time.Second
				timer := time.NewTimer(intervalTime)
				<-timer.C

				go updateStockPrice(url, selector, ticker, intervalTime)

				if len(tickers) == index+1 {
					go readyToCrawl()
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
