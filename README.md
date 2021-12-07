# Graphql server to get stock values
This is a demo project about how to get stock quotes using a web crawler with Go.

## Clone
```
git clone git@github.com:boniattirodrigo/stock.git
```

## Run
Make sure that you're using the right Go version defined on .tool-versions.

You need to set a Postgres database url, copying and changing the .env file:
```
cp .env-sample .env
```
Create the database that you defined on the database url.

Spin it up with:
```
go run main.go
```

Access http://localhost:8080/ to interact with GraphQL Playground. Here some examples to run:

**subscription:**
```
subscription {
  stocks(tickers:["PETR4", "LREN3", "AZUL4"]) {
    ticker
    price
  }
}
```

**query:**
```
query {
  stocks(tickers:["PETR4", "LREN3", "AZUL4"]) {
    ticker
    price
  }
}
```


## Build
```
go build -o bin/stock
```

## Deploy
```
git push heroku main:master
```

## Format code
```
gofmt -w -s -d .
```
