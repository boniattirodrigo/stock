package ws

import (
	"encoding/json"
	"github.com/boniattirodrigo/stock/graphql/schema"
	"github.com/gorilla/websocket"
	"github.com/graphql-go/graphql"
	"log"
)

func StockPublisher() {
	Subscribers.Range(func(key, value interface{}) bool {
		subscriber, ok := value.(*Subscriber)
		if !ok {
			return true
		}
		payload := graphql.Do(graphql.Params{
			Schema:        schema.Schema,
			RequestString: subscriber.RequestString,
		})
		message, err := json.Marshal(map[string]interface{}{
			"type":    "data",
			"id":      subscriber.OperationID,
			"payload": payload,
		})
		if err != nil {
			log.Printf("failed to marshal message: %v", err)
			return true
		}
		if err := subscriber.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			if err == websocket.ErrCloseSent {
				Subscribers.Delete(key)
				return true
			}
			log.Printf("failed to write to ws connection: %v", err)
			return true
		}
		return true
	})
}
