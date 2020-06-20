package utils

func InterfaceToArrayOfStrings(tickers interface{}) ([]string) {
  tickerInterface, _ := tickers.([]interface{})
  tickersList := make([]string, len(tickerInterface))
  for i := range tickersList {
    tickersList[i] = tickerInterface[i].(string)
  }

  return tickersList
}
