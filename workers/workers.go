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
	ws.StockChangedPublish(ticker)
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
				ws.StockChangedPublish(ticker)
			}
		}
	})

	c.Visit(fmt.Sprint(url, ticker))
}

func isTimeToCrawl() bool {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	hour := time.Now().In(loc).Hour()
	weekday := int(time.Now().In(loc).Weekday())
	saturday := 6
	sunday := 0

	return hour >= 10 && hour <= 16 && weekday != saturday && weekday != sunday
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

				if isTimeToCrawl() {
					go updateStockPrice(url, selector, ticker, intervalTime)
				}

				if len(tickers) == index+1 {
					go readyToCrawl()
				}
			}
		}
	}
}

func pingAllSubscribers() {
	// Keep pinging all subscribers every 50 seconds to avoid a timeout from Heroku
	timeTicker := time.NewTicker(45 * time.Second)

	for {
		select {
		case <-timeTicker.C:
			go ws.PingAllSubscribers()
		}
	}
}

func Start() {
	var stocks []models.Stock
	var tickersAsc []string
	var tickersDesc []string
	var funds []string
	var cryptos []string
	db.Connection.Where("type = ?", "stock").Order("ticker asc").Find(&stocks).Pluck("ticker", &tickersAsc)
	db.Connection.Where("type = ?", "stock").Order("ticker desc").Find(&stocks).Pluck("ticker", &tickersDesc)
	db.Connection.Where("type = ?", "fund").Order("ticker desc").Find(&stocks).Pluck("ticker", &funds)
	db.Connection.Where("type = ?", "crypto").Order("ticker desc").Find(&stocks).Pluck("ticker", &cryptos)

	go crawlPages("https://statusinvest.com.br/acoes/", ".special strong", tickersAsc, 5)
	go crawlPages("https://www.infomoney.com.br/", ".line-info .value p", tickersDesc, 5)
	go crawlPages("https://www.infomoney.com.br/", ".line-info .value p", funds, 10)
	go crawlPages("https://cryptowat.ch/assets/", ".items-center .price", cryptos, 10)
	go pingAllSubscribers()
}
