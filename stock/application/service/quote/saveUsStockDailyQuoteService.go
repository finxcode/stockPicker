package quote

import (
	"stockPicker/stock/application/port/out"
	"stockPicker/stock/application/service/metadata"
	"stockPicker/stock/domain/entity"
)

type saveUsStockDailyQuoteService struct {
	getUsStockMetaDataPort   out.GetUsStockMetaDataPort
	getUsStockDailyQuotePort out.GetUsStockDailyQuotePort
}

func NewSaveUsStockDailyQuoteService(getUsStockMetaDataPort out.GetUsStockMetaDataPort,
	getUsStockDailyQuotePort out.GetUsStockDailyQuotePort) *saveUsStockDailyQuoteService {
	return &saveUsStockDailyQuoteService{
		getUsStockMetaDataPort:   getUsStockMetaDataPort,
		getUsStockDailyQuotePort: getUsStockDailyQuotePort,
	}
}

func (s *saveUsStockDailyQuoteService) getUsStockMetaData() *[]entity.UsStock {
	var usCommonStock []entity.UsStock
	stocks := s.getUsStockMetaDataPort.GetUsStockMetaData()
	for _, stock := range *stocks {
		if (stock.Mic == metadata.XNYS || stock.Mic == metadata.XNAS || stock.Mic == metadata.XASE) && stock.EquityType == metadata.Type {
			usCommonStock = append(usCommonStock, stock)
		}
	}
	return &usCommonStock
}
