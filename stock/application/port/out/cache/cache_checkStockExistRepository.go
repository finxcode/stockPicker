package cache

import "github.com/google/uuid"

type CheckStockExistRepository interface {
	CheckStockExistInCache(stockId uuid.UUID) bool
}
