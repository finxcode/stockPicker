package repository

import (
	"stockPicker/stock/domain/entity"
)

type StockDataFetcher interface {
	Fetch(source string) *entity.DailyStockPrice
}
