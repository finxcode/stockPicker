package in

import (
	"fmt"
	"go.uber.org/zap"
	"stockPicker/stock/application/port/in"
	"stockPicker/stock/global"
)

type saveUsStockDailyQuoteController struct {
	saveUsStockDailyQuoteUseCase in.SaveUsStockDailyQuoteUseCase
}

func NewSaveUsStockDailyQuoteController(saveUsStockDailyQuoteUseCase in.SaveUsStockDailyQuoteUseCase) *saveUsStockDailyQuoteController {
	return &saveUsStockDailyQuoteController{
		saveUsStockDailyQuoteUseCase: saveUsStockDailyQuoteUseCase,
	}
}

func (s *saveUsStockDailyQuoteController) SaveUsStockDailyQuote() {
	numSaved := s.saveUsStockDailyQuoteUseCase.SaveUsStockDailyQuotes()
	global.App.Logger.Info("stock quote saved", zap.String("stock quote saved",
		fmt.Sprintf("%d stock quotes saved", numSaved)))
	fmt.Sprintf("%d stock quotes saved\n", numSaved)
}
