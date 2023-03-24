package out

import "stockPicker/stock/domain/entity"

type saveUsStockMetaDataConsoleController struct {
}

func NewSaveUsStockMetaDataConsoleController() *saveUsStockMetaDataConsoleController {
	return &saveUsStockMetaDataConsoleController{}
}

func (s *saveUsStockMetaDataConsoleController) SaveUsStockMetaDataInCache(stock *entity.UsStock) bool {
	return false
}

func (s *saveUsStockMetaDataConsoleController) SaveUsStockMetaDataInDB(stock *entity.UsStock) bool {
	return false
}
