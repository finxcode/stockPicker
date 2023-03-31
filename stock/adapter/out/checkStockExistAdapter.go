package out

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
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
	cmd := c.rds.Get(context.Background(), figi)
	if cmd.Err() == redis.Nil {
		return false
	} else if cmd.Err() != nil {
		global.App.Logger.Error("redis error", zap.String("redis query error",
			fmt.Sprintf("query key %s failed with error %s", figi, cmd.Err().Error())))
		return true
	}
	return true
}

func (c *checkStockExistAdapter) CheckStockExistInDB(figi string) bool {
	var id uuid.UUID
	query := "select stock_id from us_stock_meta_data where figi = ?"
	if err := c.db.Get(&id, query, figi); err == sql.ErrNoRows {
		return false
	} else if err != nil {
		global.App.Logger.Error("db error", zap.String("database query error",
			fmt.Sprintf("query figi：%s failed with error %s", figi, err.Error())))
		return true
	}
	return true
}
