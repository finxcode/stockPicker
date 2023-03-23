package db

import "github.com/google/uuid"

type CheckStockExistRepository interface {
	CheckStockExistInDB(stockId uuid.UUID) bool
}
