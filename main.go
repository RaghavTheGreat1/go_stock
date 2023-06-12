package main

import (
	"fmt"
)

func main() {
	var stockSymbol string
	fmt.Print("Enter the stock symbol: ")
	fmt.Scanf("%s", &stockSymbol)
	stockData := GetStockData(stockSymbol)
	fmt.Printf("%+v\n", stockData)
}
