package repository

import (
	"github.com/google/uuid"
	"stockPicker/internal/domain/stock"
)

type StockRepository interface {
	Add(stock stock.UsStock) error
	Update(stock stock.UsStock) error
	Delete(stockId uuid.UUID) error
	GetById(stockId uuid.UUID) (stock.UsStock, error)
	GetBySymbol(symbol string) (stock.UsStock, error)
}
