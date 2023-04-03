package main

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"stockPicker/ext/finnhub/fetcher"
	"stockPicker/ext/xueqiu/parser"
	"stockPicker/stock/adapter/in"
	"stockPicker/stock/adapter/out"
	"stockPicker/stock/application/service/metadata"
	"stockPicker/stock/application/service/quote"
	"stockPicker/stock/global"
	"stockPicker/stock/init/config"
	logging "stockPicker/stock/init/log"
	"stockPicker/stock/init/mysql"
	"stockPicker/stock/init/redis"
)

func main() {
	err, c := config.New()
	if err != nil {
		log.Fatal(fmt.Sprintf("read in configurations failed with error %s", err.Error()))
	}

	global.App.Logger = logging.New(c)
	logger := global.App.Logger

	db, err := mysql.InitDb(c)

	defer func() {
		err := db.Close()
		if err != nil {
			logger.Fatal("mysql connection close failed with error", zap.String("fatal", err.Error()))
		}
	}()

	rds, err := redis.InitRedis(c)
	if err != nil {
		logger.Fatal("redis initialization failed with error", zap.String("redis fatal error", err.Error()))
	}

	defer func() {
		err := rds.Close()
		if err != nil {
			logger.Fatal("redis connection close failed with error", zap.String("redis fatal error", err.Error()))
		}

	}()

	StockSymbolFetcher := fetcher.NewStockSymbolFetcher()
	getUsStockMetaDataAdapter := out.NewGetUsStockMetaDataAdapter(StockSymbolFetcher)
	checkStockExistAdapter := out.NewCheckStockExistAdapter(rds, db)
	saveUsStockMetaDataAdapter := out.NewSaveUsStockMetaDataConsoleController(rds, db)
	getUsStocksAdapter := out.NewGetUsStocksAdapter(rds, db)
	usStockMetaDataService := metadata.NewUsStockSymbolService(
		getUsStockMetaDataAdapter, checkStockExistAdapter,
		checkStockExistAdapter, saveUsStockMetaDataAdapter,
		saveUsStockMetaDataAdapter, getUsStocksAdapter)
	saveUsStockConsoleController := in.NewSaveUsStockConsoleController(usStockMetaDataService)

	_ = saveUsStockConsoleController.SaveUsStockMetaData()

	stockQuoteParser := parser.NewStockPriceParser()
	getUsStockDailyQuoteAdapter := out.NewGetUsStockDailyQuoteAdapter(stockQuoteParser)
	getUsStockAdapter := out.NewGetUsStocksAdapter(rds, db)
	saveUsStockDailyQuoteAdapter := out.NewSaveUsStockDailyQuoteAdapter(db)
	saveUsStockDailyQuoteService := quote.NewSaveUsStockDailyQuoteService(c, getUsStockDailyQuoteAdapter,
		getUsStockAdapter, getUsStockAdapter, saveUsStockDailyQuoteAdapter)
	saveUsStockDailyQuoteController := in.NewSaveUsStockDailyQuoteController(saveUsStockDailyQuoteService)
	saveUsStockDailyQuoteController.SaveUsStockDailyQuote()

}
