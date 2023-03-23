package main

import (
	"fmt"
	"stockPicker/ext/finnhub/fetcher"
)

func main() {

	f := fetcher.NewStockSymbolFetcher()
	s, err := f.GetUsStockSymbol()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(len(*s))
}
