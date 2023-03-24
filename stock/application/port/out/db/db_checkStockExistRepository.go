package db

type CheckStockExistRepository interface {
	CheckStockExistInDB(figi string) bool
}
