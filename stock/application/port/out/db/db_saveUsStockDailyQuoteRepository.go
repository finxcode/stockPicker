package db

import "stockPicker/stock/domain/entity"

type SaveUsStockDailyQuoteRepository interface {
	SaveUsStockDailyQuoteToDb(quote *entity.StockDailyQuote) error
}
