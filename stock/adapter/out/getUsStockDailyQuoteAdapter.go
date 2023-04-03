package out

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"stockPicker/ext/xueqiu/parser"
	"stockPicker/stock/adapter/out/mapper"
	"stockPicker/stock/adapter/out/model"
	"stockPicker/stock/domain/entity"
	"stockPicker/stock/global"
)

type getUsStockDailyQuoteAdapter struct {
	stockQuoteParser *parser.StockQuoteParser
}

func NewGetUsStockDailyQuoteAdapter(stockQuoteParser *parser.StockQuoteParser) *getUsStockDailyQuoteAdapter {
	return &getUsStockDailyQuoteAdapter{
		stockQuoteParser: stockQuoteParser,
	}
}

func (g *getUsStockDailyQuoteAdapter) GetUsStockDailyQuote(url, symbol string) (*entity.StockDailyQuote, error) {
	quotes, err := g.stockQuoteParser.ParseStockData(url)
	if err != nil {
		global.App.Logger.Error("parse stock data error", zap.String("parse stock data failed",
			fmt.Sprintf("parse url: %s faile with error: %s", url, err.Error())))
		return nil, err
	}

	var quoteDataEntity model.UsStockDailyQuote

	config := &mapstructure.DecoderConfig{
		ErrorUnused: true,
		Result:      &quoteDataEntity,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}

	if err := decoder.Decode(quotes); err != nil {
		return nil, err
	}

	quoteDataEntity.SetTradingDay(int64(quoteDataEntity.FetchTime))
	quoteDataEntity.SetSymbol(symbol)

	return mapper.MapQuoteDataEntityToQuoteDomainEntity(&quoteDataEntity), nil
}
