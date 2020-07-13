package migrations

import (
	"encoding/json"
	"fmt"
	"github.com/boniattirodrigo/stock/models"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"os"
)

type Ticker struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Tickers struct {
	Tickers []Ticker `json:"tickers"`
}

func readTickersFromJsonFile() []Ticker {
	jsonFile, err := os.Open("tickers.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var tickers Tickers

	json.Unmarshal(byteValue, &tickers)

	return tickers.Tickers
}

func CreateStocks(dbConnection *gorm.DB) {
	tickers := readTickersFromJsonFile()

	dbConnection.AutoMigrate(&models.Stock{})

	for _, ticker := range tickers {
		var stock models.Stock
		dbConnection.FirstOrCreate(&stock, models.Stock{Ticker: ticker.Name, Type: ticker.Type})
	}
}
