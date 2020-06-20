package main

import (
  "fmt"
  "time"
  "github.com/boniattirodrigo/stock/db"
  "github.com/boniattirodrigo/stock/services"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var exit = make(chan bool)

func main() {
  stocks := [3]string{"AZUL4", "PETR4", "LREN3"}
  dbConnection, err := gorm.Open("postgres", "user=user dbname=dbname sslmode=disable")

  if err != nil {
    panic(err)
  }

  defer dbConnection.Close()

  db.RunMigrations(stocks, dbConnection)

  ticker := time.NewTicker(5 * time.Second)
  go func() {
    for {
        select {
        case <-ticker.C:
          for _, stockName := range stocks {
            go services.UpdateStocks(stockName, dbConnection)
          }
        }
    }
  }()

  <-exit
  fmt.Println("Done.")
}
