package out

import "stockPicker/stock/domain/entity"

type GetUsStockDailyQuotePort interface {
	GetUsStockDailyQuote(url, symbol string) (*entity.StockDailyQuote, error)
}
