package ws

import (
	"encoding/json"
	"github.com/boniattirodrigo/stock/graphql/schema"
	"github.com/boniattirodrigo/stock/utils"
	"github.com/gorilla/websocket"
	"github.com/graphql-go/graphql"
	"log"
)

func publish(subscriber *Subscriber, key int) bool {
	payload := graphql.Do(graphql.Params{
		Schema:         schema.Schema,
		RequestString:  subscriber.RequestString,
		VariableValues: subscriber.Variables,
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
}

func InitialPublish(key int) {
	subscriber, ok := Subscribers.Load(key)

	if ok {
		publish(subscriber.(*Subscriber), key)
	}
}

func PingAllSubscribers() {
	Subscribers.Range(func(key, value interface{}) bool {
		subscriber, ok := value.(*Subscriber)

		if ok {
			publish(subscriber, key.(int))
		}

		return true
	})
}

func StockChangedPublish(ticker string) {
	Subscribers.Range(func(key, value interface{}) bool {
		subscriber, ok := value.(*Subscriber)
		subscriberTickers := utils.InterfaceToArrayOfStrings(subscriber.Variables["tickers"])
		tickerFound := false

		for _, currentTicker := range subscriberTickers {
			if currentTicker == ticker {
				tickerFound = true
				break
			}
		}

		if ok && tickerFound {
			publish(subscriber, key.(int))
		}

		return true
	})
}
