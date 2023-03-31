package db

import "stockPicker/stock/domain/entity"

type GetUsStockByFigiRepository interface {
	GetUsStockByFigi(figi string) (*entity.UsStock, error)
}
