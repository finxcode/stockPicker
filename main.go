package main

import (
	"fmt"
	"stockPicker/ext/xueqiu/parser"
	"stockPicker/stock/adapter/out"
)

func main() {
	//var cstSh, _ = time.LoadLocation("EST") //上海
	//fmt.Println(time.Unix(1680120000526/1000, 0).In(cstSh).Format("2006-01-02 15:04:05"))
	//fmt.Println(time.Unix(0, 1680120000526*int64(time.Millisecond)).In(cstSh).Format("2006-01-02 15:04:05"))

	p := parser.NewStockPriceParser()
	getUsStockDailyQuoteAdapter := out.NewGetUsStockDailyQuoteAdapter(p)
	quote, err := getUsStockDailyQuoteAdapter.GetUsStockDailyQuote("https://xueqiu.com/S/LKCO")

	if err != nil {
		fmt.Println(err.Error())
	} else {

		fmt.Println(quote)
	}
}
