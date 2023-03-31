package mapper

import (
	"stockPicker/stock/adapter/out/model"
	"stockPicker/stock/domain/entity"
)

func MapQuoteDataEntityToQuoteDomainEntity(quoteDataEntity *model.UsStockDailyQuote) *entity.StockDailyQuote {
	return entity.NewStockDailyQuote(
		quoteDataEntity.Symbol, quoteDataEntity.Current, quoteDataEntity.LastClose, quoteDataEntity.High, quoteDataEntity.Low,
		quoteDataEntity.Open, quoteDataEntity.Change, quoteDataEntity.Percent, quoteDataEntity.High52Week, quoteDataEntity.Low52Week,
		quoteDataEntity.Volume, quoteDataEntity.Amount, quoteDataEntity.MarketCapital, quoteDataEntity.TotalShares,
		quoteDataEntity.Dividend, quoteDataEntity.DividendYield, quoteDataEntity.EPS, quoteDataEntity.TurnoverRate, quoteDataEntity.TurnoverRate,
		quoteDataEntity.Amplitude, quoteDataEntity.CurrentYearPercent, quoteDataEntity.IssueDate, quoteDataEntity.PeTtm,
		quoteDataEntity.PeLyr, quoteDataEntity.PeForecast, quoteDataEntity.Navps, quoteDataEntity.Pb, quoteDataEntity.Psr,
		quoteDataEntity.Timestamp, quoteDataEntity.TradingDay)
}
