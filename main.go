package main

import (
	"fmt"
	"flag"
	"github.com/gocolly/colly"
)

func main() {
	htmlSelector := flag.String("selector", "", "HTML selector to find the data")
	url := flag.String("url", "", "URL to fetch the information")
	stock := flag.String("stock", "", "Stock that you want to know the value")
	flag.Parse()

	c := colly.NewCollector()

	c.OnHTML(*htmlSelector, func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.Visit(fmt.Sprint(*url, *stock))
}
