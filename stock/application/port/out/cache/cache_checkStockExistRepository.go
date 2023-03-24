package cache

type CheckStockExistRepository interface {
	CheckStockExistInCache(figi string) bool
}
