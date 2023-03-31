package cache

type GetUsStocksByFigiRepository interface {
	GetUsStocksByFigi() (figi, symbol string)
}
