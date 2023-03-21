package repository

import (
	"github.com/google/uuid"
	"stockPicker/stock/domain/entity"
)

type StockMetaDataRepository interface {
	Add(stock entity.UsStock) error
	Update(stock entity.UsStock) error
	Delete(stockId uuid.UUID) error
	GetById(stockId uuid.UUID) (entity.UsStock, error)
	GetBySymbol(symbol string) (entity.UsStock, error)
}
