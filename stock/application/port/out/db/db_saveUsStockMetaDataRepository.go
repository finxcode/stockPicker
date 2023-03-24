package db

import "stockPicker/stock/domain/entity"

type SaveUsStockMetaDataRepository interface {
	SaveUsStockMetaDataInDB(stocks *entity.UsStock) bool
}
