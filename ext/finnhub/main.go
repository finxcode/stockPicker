package main

import (
	"fmt"
	"stockPicker/ext/finnhub/config"
)

func main() {
	err, c := config.New()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(c.FinnHub.Stock)
}
