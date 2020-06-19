# Get stock value by CLI

This is a demo project about how to get a stock value using a web crawler.

## Clone
```
git clone git@github.com:boniattirodrigo/stock.git
```

## Build
Make sure that you are using the right Go version defined on .tool-versions
```
go build
```

## Run
To run you need to pass 3 parameters:
- selector = HTML selector to find the data
- url = URL to fetch the information
- stock = Stock that you want to know the value

You can also run `./stock -h` to see these parameters.

**Example:**
```
./stock -selector=.stock-value -url=https://stock-values.com/stock/ -stock=stk3
```

You can also pass an array of stocks:
```
./stock -selector=.stock-value -url=https://stock-values.com/stock/ -stock=stk3,stk4,cmp3
```
