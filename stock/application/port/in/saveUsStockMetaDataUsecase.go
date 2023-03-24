package in

import "stockPicker/stock/domain/entity"

type SaveUsStockMetaDataUseCase interface {
	SaveUsStockMetaData(*[]entity.UsStock) (int, int)
}
