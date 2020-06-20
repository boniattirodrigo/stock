package types

import "github.com/graphql-go/graphql"

var StockType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Stock",
	Description: "Stock market",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The id of the stock.",
		},
		"ticker": &graphql.Field{
			Type:        graphql.String,
			Description: "Stock ticker",
    },
		"price": &graphql.Field{
			Type:        graphql.Float,
			Description: "Stock price",
    },
		"createdAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "Stock created date",
    },
		"updatedAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "Stock updated date",
		},
	},
})
