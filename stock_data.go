package main

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

type StockData struct {
	Symbol    string
	Open      string
	High      string
	Low       string
	Close     string
	Volume    string
	TimeStamp string
}

func GetStockData(symbol string) StockData {
	var query string = fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%s&interval=5min&apikey=%s", symbol, APIkey)
	response, err := http.Get(query)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseBody, _ := io.ReadAll(response.Body)

	//	fmt.Println(string(responseBody))

	var rawStockData map[string]map[string]any

	json.Unmarshal(responseBody, &rawStockData)

	var parsedStockData StockData

	timeSeriesData := rawStockData["Time Series (5min)"]

	for key, value := range timeSeriesData {
		dateTimeMap := value.(map[string]any)

		parsedStockData = StockData{
			Symbol:    symbol,
			Open:      dateTimeMap["1. open"].(string),
			High:      dateTimeMap["2. high"].(string),
			Low:       dateTimeMap["3. low"].(string),
			Close:     dateTimeMap["4. close"].(string),
			Volume:    dateTimeMap["5. volume"].(string),
			TimeStamp: key,
		}

		break

	}

	return parsedStockData
}
