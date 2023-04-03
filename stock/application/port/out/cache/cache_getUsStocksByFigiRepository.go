package cache

type GetUsStocksByFigiRepository interface {
	GetUsStockSymbolByFigi(figi string) string
}
