package out

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
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

func (c *checkStockExistAdapter) CheckStockExistInCache(stockId uuid.UUID) bool {
	return false
}

func (c *checkStockExistAdapter) CheckStockExistInDB(stockId uuid.UUID) bool {
	return false
}
