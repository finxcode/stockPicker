package out

import (
	"stockPicker/ext/finnhub/fetcher"
	"stockPicker/stock/adapter/out/mapper"
	"stockPicker/stock/domain/entity"
)

type getUsStockMetaDataAdapter struct {
	stockSymbolFetcher *fetcher.StockSymbolFetcher
}

func NewGetUsStockMetaDataAdapter(symbolFetcher *fetcher.StockSymbolFetcher) *getUsStockMetaDataAdapter {
	return &getUsStockMetaDataAdapter{
		stockSymbolFetcher: symbolFetcher,
	}
}

func (g *getUsStockMetaDataAdapter) GetUsStockMetaData() *[]entity.UsStock {
	var usStocks []entity.UsStock
	finnHubSymbol, err := g.stockSymbolFetcher.GetUsStockSymbol()

	if err != nil {
		return nil
	} else {
		for _, symbol := range *finnHubSymbol {
			usStocks = append(usStocks, *mapper.StockSymbolToStockMeta(symbol))
		}
	}
	return &usStocks
}
