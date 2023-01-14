package repository

import "stockPicker/internal/domain/stock"

type FetcherRepository interface {
	FetchStockSymbol() (*[]stock.Symbol, error)
}
