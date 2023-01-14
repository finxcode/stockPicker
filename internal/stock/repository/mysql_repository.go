package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"stockPicker/internal/domain/stock"
	"stockPicker/internal/repository"
)

type mysqlStockRepository struct {
	Db *sqlx.DB
}

func NewMysqlStockRepository(db *sqlx.DB) repository.StockRepository {
	return &mysqlStockRepository{
		Db: db,
	}
}

func (m *mysqlStockRepository) Add(stock stock.UsStock) error {
	//TODO implement me
	panic("implement me")
}

func (m *mysqlStockRepository) Update(stock stock.UsStock) error {
	//TODO implement me
	panic("implement me")
}

func (m *mysqlStockRepository) Delete(stockId uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (m *mysqlStockRepository) GetById(stockId uuid.UUID) (stock.UsStock, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mysqlStockRepository) GetBySymbol(symbol string) (stock.UsStock, error) {
	//TODO implement me
	panic("implement me")
}
