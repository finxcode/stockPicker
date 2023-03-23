package in

import "stockPicker/stock/application/port/in"

type saveUsStockConsoleController struct {
	saveUsStockMetaDataUseCase in.SaveUsStockMetaDataUseCase
}

func NewSaveUsStockConsoleController(saveUsStockMetaDataUseCase in.SaveUsStockMetaDataUseCase) *saveUsStockConsoleController {
	return &saveUsStockConsoleController{
		saveUsStockMetaDataUseCase: saveUsStockMetaDataUseCase,
	}
}

func (s *saveUsStockConsoleController) SaveUsStockMetaData() error {
	return nil
}
