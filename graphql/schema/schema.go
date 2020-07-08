package schema

import (
	"github.com/boniattirodrigo/stock/graphql/resolvers"
	"github.com/boniattirodrigo/stock/graphql/types"
	"github.com/boniattirodrigo/stock/utils"
	"github.com/graphql-go/graphql"
)

var Schema graphql.Schema

func init() {
	fields := graphql.Fields{
		"stocks": &graphql.Field{
			Type: graphql.NewList(types.StockType),
			Args: graphql.FieldConfigArgument{
				"tickers": &graphql.ArgumentConfig{
					Description: "Tickers",
					Type:        graphql.NewNonNull(graphql.NewList(graphql.String)),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				tickers := utils.InterfaceToArrayOfStrings(p.Args["tickers"])
				return resolvers.FetchStocks(tickers), nil
			},
		},
	}

	Query := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Query",
		Fields: fields,
	})

	Subscription := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Subscription",
		Fields: fields,
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:        Query,
		Subscription: Subscription,
	})
	if err != nil {
		panic(err)
	}
	Schema = schema
}
