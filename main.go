package main

import (
  "fmt"
  "net/http"
  "github.com/boniattirodrigo/stock/workers"
  "github.com/boniattirodrigo/stock/ws"
  "github.com/boniattirodrigo/stock/graphql/schema"
  "github.com/boniattirodrigo/stock/db"
  "github.com/graphql-go/handler"
)

func main() {
  db.Connect()
  go workers.Start()

	h := handler.New(&handler.Config{
		Schema:     &schema.Schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	http.Handle("/", h)

	http.HandleFunc("/subscriptions", ws.Handler(ws.StockPublisher))

	fmt.Println("server is started at: http://localhost:8080/")
	fmt.Println("graphql api server is started at: http://localhost:8080/graphql")
	fmt.Println("subscriptions api server is started at: http://localhost:8080/subscriptions")
  http.ListenAndServe(":8080", nil)
}
