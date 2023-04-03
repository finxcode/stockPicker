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
)

type getUsStockAdapter struct {
	rds *redis.Client
	db  *sqlx.DB
}

func NewGetUsStocksAdapter(rds *redis.Client, db *sqlx.DB) *getUsStockAdapter {
	return &getUsStockAdapter{
		rds: rds,
		db:  db,
	}
}

//GetUsStockByFigi
//get stock in db
func (g *getUsStockAdapter) GetUsStockByFigi(figi string) (*entity.UsStock, error) {
	var stock model.UsStock
	query := "SELECT * from us_stock_meta_data where figi = ?"
	err := g.db.QueryRowx(query, figi).Scan(&stock)
	if err != nil {
		global.App.Logger.Error("database error", zap.String("query us_stock_meta_data error",
			fmt.Sprintf("query stock figi:%s failed with error:%s", figi, err.Error())))
		return nil, err
	}
	return stock.StockDataEntityToDomainEntity(), nil
}

func (g *getUsStockAdapter) GetAllUsStockFigi() *[]string {
	var cursor uint64
	keys, cursor, err := g.rds.Scan(context.Background(), cursor, "*", 0).Result()
	if err != nil {
		global.App.Logger.Error("redis error", zap.String("redis error getting all keys", err.Error()))
		return nil
	}
	if cursor == 0 {
		global.App.Logger.Info("redis no key found", zap.String("redis no key found", "no key found"))
		return nil
	}

	return &keys
}

func (g *getUsStockAdapter) GetUsStockSymbolByFigi(figi string) string {
	res, err := g.rds.Get(context.Background(), figi).Result()
	if err != nil {
		global.App.Logger.Error("redis error", zap.String("redis error when getting key",
			fmt.Sprintf("failed to get key: %s with error：%s", figi, err.Error())))
		return ""
	}
	return res
}
