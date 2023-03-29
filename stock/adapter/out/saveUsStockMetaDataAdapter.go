package out

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
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
	cmd := s.rds.Set(context.Background(), stock.Figi, stock.Symbol, time.Hour*24*30)
	if cmd.Err() != nil {
		global.App.Logger.Error("redis error", zap.String("redis set key error",
			fmt.Sprintf("redis set key:%s failed with error:%s", stock.Figi, cmd.Err().Error())))
		return false
	}
	return true
}

func (s *saveUsStockMetaDataConsoleController) SaveUsStockMetaDataInDB(domainStock *entity.UsStock) bool {
	//stock := dataEntity.NewUsStockDataEntity(uuid.New(), domainStock.Currency,
	//	domainStock.Description, domainStock.DisplaySymbol,
	//	domainStock.Figi, domainStock.IsIn, domainStock.Mic,
	//	domainStock.ShareClassFigi, domainStock.Symbol, domainStock.Symbol2, domainStock.EquityType)
	//query := "INSERT INTO us-stock-meta-data "
	return false
}
