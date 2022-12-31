package main

import (
	"fmt"
	"go.uber.org/zap"
	"stockPicker/config"
	"stockPicker/global"
	"stockPicker/internal/data"
	"stockPicker/internal/log"
)

func main() {

	err, c := config.New()
	if err != nil {
		panic(fmt.Sprintf("config initialization failed with error: %s", err.Error()))
	}
	global.App.Config = c
	global.App.Logger = log.New(global.App.Config)
	global.App.Logger.Info("initialization", zap.String("logger", "started"))

	symbols, err := data.FetchStockSymbol()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(len(*symbols))

}
