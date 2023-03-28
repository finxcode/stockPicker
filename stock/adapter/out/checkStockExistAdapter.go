package out

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"stockPicker/stock/global"
)

type checkStockExistAdapter struct {
	rds *redis.Client
	db  *sqlx.DB
}

func NewCheckStockExistAdapter(rds *redis.Client, db *sqlx.DB) *checkStockExistAdapter {
	return &checkStockExistAdapter{
		rds: rds,
		db:  db,
	}
}

func (c *checkStockExistAdapter) CheckStockExistInCache(figi string) bool {
	b, err := c.rds.Get(context.Background(), figi).Bool()
	if err != nil {
		global.App.Logger.Error("redis error", zap.String("redis query error",
			fmt.Sprintf("query key %s failed with error %s", figi, err.Error())))
	}
	return b
}

func (c *checkStockExistAdapter) CheckStockExistInDB(figi string) bool {
	return false
}
