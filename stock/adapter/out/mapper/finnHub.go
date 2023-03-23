package mapper

import (
	finnHub "stockPicker/ext/finnhub/entity"
	"stockPicker/stock/domain/entity"
)

func StockSymbolToStockMeta(s finnHub.Symbol) *entity.UsStock {
	return entity.NewUsStock(
		s.Currency, s.Description, s.DisplaySymbol, s.Figi, s.Isin, s.Mic, s.ShareClassFigi, s.Symbol, s.Symbol2, s.Type,
	)
}
