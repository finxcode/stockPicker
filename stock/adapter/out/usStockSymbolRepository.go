package out

import "stockPicker/stock/domain/entity"

type UsStockSymbolRepository interface {
	SaveAll(*[]entity.UsStock) (int, error)
}
