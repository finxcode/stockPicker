package mapper

import (
	data "stockPicker/stock/adapter/out/entity"
	"stockPicker/stock/domain/entity"
)

func ConvertDomainStockToPersistenceStock(stock *entity.UsStock) *data.UsStock {
	if stock == nil {
		return nil
	}
	return &data.UsStock{
		StockId:        stock.StockId,
		Currency:       stock.Currency,
		Description:    stock.Description,
		DisplaySymbol:  stock.DisplaySymbol,
		Figi:           stock.Figi,
		IsIn:           stock.IsIn,
		Mic:            stock.Mic,
		ShareClassFigi: stock.ShareClassFigi,
		Symbol:         stock.Symbol,
		Symbol2:        stock.Symbol2,
		EquityType:     stock.EquityType,
	}
}
