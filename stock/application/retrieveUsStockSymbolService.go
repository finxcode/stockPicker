package application

import (
	"stockPicker/config"
	"stockPicker/stock/domain/entity"
)

type RetrieveUsStockSymbolService struct {
	config *config.Config
}

func NewRetrieveUsStockSymbolService(config *config.Config) *RetrieveUsStockSymbolService {
	return &RetrieveUsStockSymbolService{
		config: config,
	}
}

func (r *RetrieveUsStockSymbolService) RetrieveUsStockSymbol() *[]entity.UsStock {
	return nil
}
