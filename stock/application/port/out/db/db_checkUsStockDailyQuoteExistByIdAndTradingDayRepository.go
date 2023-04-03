package db

import (
	"github.com/google/uuid"
)

type CheckUsStockDailyQuoteExistByIdAndTradingDayRepository interface {
	CheckUsStockDailyQuoteExistByIdAndTradingDayInDb(stockId uuid.UUID, tradingDay string) bool
}
