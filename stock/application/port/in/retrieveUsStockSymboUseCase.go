package in

import "stockPicker/stock/domain/entity"

type RetrieveUsStockSymbolUseCase interface {
	RetrieveUsStockSymbol() *[]entity.UsStock
}
