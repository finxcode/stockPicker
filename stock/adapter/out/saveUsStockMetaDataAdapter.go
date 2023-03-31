package out

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"stockPicker/stock/adapter/out/model"
	"stockPicker/stock/domain/entity"
	"stockPicker/stock/global"
	"time"
)

type saveUsStockMetaDataConsoleController struct {
	rds *redis.Client
	db  *sqlx.DB
}

func NewSaveUsStockMetaDataConsoleController(rds *redis.Client, db *sqlx.DB) *saveUsStockMetaDataConsoleController {
	return &saveUsStockMetaDataConsoleController{
		rds: rds,
		db:  db,
	}
}

func (s *saveUsStockMetaDataConsoleController) SaveUsStockMetaDataInCache(stock *entity.UsStock) bool {
	cmd := s.rds.Set(context.Background(), stock.Figi,
		fmt.Sprintf("%s %s", stock.Symbol, stock.StockId), time.Hour*24*30)
	if cmd.Err() != nil {
		global.App.Logger.Error("redis error", zap.String("redis set key error",
			fmt.Sprintf("redis set key:%s failed with error:%s", stock.Figi, cmd.Err().Error())))
		return false
	}
	return true
}

func (s *saveUsStockMetaDataConsoleController) SaveUsStockMetaDataInDB(domainStock *entity.UsStock) bool {
	stock := model.NewUsStockDataEntity(domainStock.StockId, domainStock.Currency,
		domainStock.Description, domainStock.DisplaySymbol,
		domainStock.Figi, domainStock.IsIn, domainStock.Mic,
		domainStock.ShareClassFigi, domainStock.Symbol, domainStock.Symbol2, domainStock.EquityType)
	query := "INSERT INTO us_stock_meta_data (stock_id,currency,description," +
		"display_symbol,figi,is_in,mic,share_class_figi,symbol,symbol2,equity_type,created_at) " +
		"VALUES (:stock_id,:currency,:description,:display_symbol,:figi,:is_in,:mic,:share_class_figi,:symbol,:symbol2,:equity_type,:created_at)"
	_, err := s.db.NamedExec(query, stock)
	if err != nil {
		global.App.Logger.Error("db insert error", zap.String("database insertion failed",
			fmt.Sprintf("insert stock figi:%s failed with error:%s", stock.GetFigi(), err.Error())))
		return false
	}
	return true
}
