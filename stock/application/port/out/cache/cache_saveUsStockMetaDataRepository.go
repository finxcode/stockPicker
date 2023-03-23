package cache

import "stockPicker/stock/domain/entity"

type SaveUsStockMetaDataRepository interface {
	SaveUsStockMetaDataInCache(stocks *[]entity.UsStock) (int, error)
}
