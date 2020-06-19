package main

import (
	"fmt"
	"flag"
	"strings"
	"github.com/gocolly/colly"
)

func main() {
	htmlSelector := flag.String("selector", "", "HTML selector to find the data")
	url := flag.String("url", "", "URL to fetch the information")
	stock := flag.String("stock", "", "Stock that you want to know the value")
	flag.Parse()

	c := colly.NewCollector()

	c.OnHTML(*htmlSelector, func(e *colly.HTMLElement) {
    stockName := strings.ReplaceAll(e.Request.URL.String(), *url, "")

		fmt.Printf("%s: %s \n", stockName, e.Text)
	})

  stocks := strings.Split(*stock, ",")

  for _, v := range stocks {
    c.Visit(fmt.Sprint(*url, v))
  }
}
