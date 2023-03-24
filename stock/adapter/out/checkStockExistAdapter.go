package out

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type checkStockExistAdapter struct {
	db *sqlx.DB
}

func NewCheckStockExistAdapter(db *sqlx.DB) *checkStockExistAdapter {
	return &checkStockExistAdapter{
		db: db,
	}
}

func (c *checkStockExistAdapter) CheckStockExistInCache(figi string) bool {
	return false
}

func (c *checkStockExistAdapter) CheckStockExistInDB(figi string) bool {
	return false
}
