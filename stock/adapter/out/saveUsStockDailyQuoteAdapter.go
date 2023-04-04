package out

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"stockPicker/stock/adapter/out/model"
	"stockPicker/stock/domain/entity"
	"stockPicker/stock/global"
)

type saveUsStockDailyQuoteAdapter struct {
	db *sqlx.DB
}

func NewSaveUsStockDailyQuoteAdapter(db *sqlx.DB) *saveUsStockDailyQuoteAdapter {
	return &saveUsStockDailyQuoteAdapter{
		db: db,
	}
}

func (s *saveUsStockDailyQuoteAdapter) SaveUsStockDailyQuoteToDb(quote *entity.StockDailyQuote) error {
	quoteDataEntity := model.NewStockDailyQuote(quote.Symbol, quote.Current, quote.LastClose, quote.High, quote.Low,
		quote.Open, quote.Change, quote.Percent, quote.High52Week, quote.Low52Week, quote.Volume, quote.Amount,
		quote.MarketCapital, quote.TotalShares, quote.Dividend, quote.DividendYield, quote.EPS, quote.TurnoverRate,
		quote.VolumeRatio, quote.Amplitude, quote.CurrentYearPercent, quote.IssueDate, quote.PeTtm, quote.PeLyr,
		quote.PeForecast, quote.Navps, quote.Pb, quote.Psr, quote.Timestamp, quote.TradingDay)
	quoteDataEntity.SetStockId(quote.StockId)
	if s.CheckUsStockDailyQuoteExistByIdAndTradingDayInDb(quoteDataEntity.StockId, quoteDataEntity.TradingDay) {
		global.App.Logger.Info("stock quote already exists", zap.String("record exist",
			fmt.Sprintf("quote stock_id: %s symbol:%s on trading_day:%s already exist",
				quoteDataEntity.StockId, quoteDataEntity.Symbol, quoteDataEntity.TradingDay)))
		return errors.New("quote already exist or error occurred in db operation")
	} else {
		query := "INSERT INTO us_stock_daily_quote (stock_id, symbol, current, last_close, high, low, open, chg," +
			"percent, high52w, low52w, volume, amount, market_capital, total_shares, dividend, dividend_yield, " +
			"eps, turnover_rate, volume_ratio, amplitude, current_year_percent, issue_date, pe_ttm, pe_lyr," +
			"pe_forecast, navps, pb,psr,fetch_time, trading_day, created_at)" +
			"VALUES (:stock_id, :symbol, :current, :last_close, :high, :low, :open, :chg," +
			":percent, :high52w, :low52w, :volume, :amount, :market_capital, :total_shares, :dividend, :dividend_yield, " +
			":eps, :turnover_rate, :volume_ratio, :amplitude, :current_year_percent, :issue_date, :pe_ttm, :pe_lyr," +
			":pe_forecast, :navps, :pb, :psr, :fetch_time, :trading_day, :created_at)"

		_, err := s.db.NamedExec(query, quoteDataEntity)
		if err != nil {
			global.App.Logger.Error("db error", zap.String("us_stock_daily_quote db insertion error",
				fmt.Sprintf("insert quote stock_id:%s symbol:%s trading_day:%s failed with error: %s",
					quoteDataEntity.StockId, quoteDataEntity.Symbol, quoteDataEntity.TradingDay, err.Error())))
			return err
		}

		global.App.Logger.Info("stock daily quote saved", zap.String("one quote saved",
			fmt.Sprintf("quote stock_id: %s symbol:%s on trading_day:%s saved sucessfully",
				quoteDataEntity.StockId, quoteDataEntity.Symbol, quoteDataEntity.TradingDay)))

		return nil
	}

}

func (s *saveUsStockDailyQuoteAdapter) CheckUsStockDailyQuoteExistByIdAndTradingDayInDb(
	stockId uuid.UUID, tradingDay string) bool {
	var quoteDataEntity model.UsStockDailyQuote
	query := "select * from us_stock_daily_quote where stock_id = ? and trading_day = ?"
	if err := s.db.Get(&quoteDataEntity, query, stockId, tradingDay); err == sql.ErrNoRows {
		return false
	} else if err != nil {
		global.App.Logger.Error("db error", zap.String("database query error",
			fmt.Sprintf("query stock_id：%s and trading day: %s failed with error %s", stockId, tradingDay, err.Error())))
		return true
	}
	return true
}
