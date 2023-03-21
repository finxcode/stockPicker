package repository

import (
	"stockPicker/stock/domain/entity"
)

type UsStockSymbolFetcherRepository interface {
	FetchUsStockSymbol() (*[]entity.UsStock, error)
}
