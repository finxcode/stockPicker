package out

import "stockPicker/stock/domain/entity"

type saveUsStockMetaDataConsoleController struct {
}

func NewSaveUsStockMetaDataConsoleController() *saveUsStockMetaDataConsoleController {
	return &saveUsStockMetaDataConsoleController{}
}

func (s *saveUsStockMetaDataConsoleController) SaveUsStockMetaDataInCache(stocks *[]entity.UsStock) (int, error) {
	return 0, nil
}

func (s *saveUsStockMetaDataConsoleController) SaveUsStockMetaDataInDB(stocks *[]entity.UsStock) (int, error) {
	return 0, nil
}
