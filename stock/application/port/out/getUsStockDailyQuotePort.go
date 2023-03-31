package out

import "stockPicker/stock/domain/entity"

type GetUsStockDailyQuotePort interface {
	GetUsStockDailyQuote(url string) (*entity.StockDailyQuote, error)
}
