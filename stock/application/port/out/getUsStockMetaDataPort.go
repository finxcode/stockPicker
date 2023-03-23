package out

import (
	"stockPicker/stock/domain/entity"
)

type GetUsStockMetaDataPort interface {
	GetUsStockMetaData() *[]entity.UsStock
}
