package main

import (
  "github.com/boniattirodrigo/stock/db"
  "github.com/boniattirodrigo/stock/db/migrations"
)

func main() {
  db.Connect()
  migrations.CreateStocks(db.Connection)
}
