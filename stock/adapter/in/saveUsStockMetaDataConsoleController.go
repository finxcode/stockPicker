package in

import (
	"fmt"
	"go.uber.org/zap"
	"stockPicker/stock/application/port/in"
	"stockPicker/stock/global"
)

type saveUsStockConsoleController struct {
	saveUsStockMetaDataUseCase in.SaveUsStockMetaDataUseCase
}

func NewSaveUsStockConsoleController(saveUsStockMetaDataUseCase in.SaveUsStockMetaDataUseCase) *saveUsStockConsoleController {
	return &saveUsStockConsoleController{
		saveUsStockMetaDataUseCase: saveUsStockMetaDataUseCase,
	}
}

func (s *saveUsStockConsoleController) SaveUsStockMetaData() error {
	numSavedInCache, numSavedInDb := s.saveUsStockMetaDataUseCase.SaveUsStockMetaData()
	if numSavedInCache != 0 || numSavedInDb != 0 {
		global.App.Logger.Info("stock meta data saved", zap.String("meta data saved",
			fmt.Sprintf("number of records saved to cache:%d number of records saved to db:%d", numSavedInCache, numSavedInDb)))
	}
	fmt.Println(fmt.Sprintf("number of records saved to cache:%d number of records saved to db:%d", numSavedInCache, numSavedInDb))
	return nil
}
