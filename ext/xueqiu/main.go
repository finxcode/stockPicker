package main

import (
	"fmt"
	"stockPicker/ext/xueqiu/parser"
)

func main() {

	p := parser.NewStockPriceParser()
	res, err := p.ParseStockData("https://xueqiu.com/S/LKCO")
	if err != nil {

	}
	fmt.Println(res)
}
