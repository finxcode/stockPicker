package model

import (
	"database/sql"
	"github.com/google/uuid"
	"stockPicker/stock/domain/entity"
	"time"
)

type UsStock struct {
	StockId        uuid.UUID    `db:"stock_id"`
	Currency       string       `db:"currency"`
	Description    string       `db:"description"`
	DisplaySymbol  string       `db:"display_symbol"`
	Figi           string       `db:"figi"`
	IsIn           string       `db:"is_in"`
	Mic            string       `db:"mic"`
	ShareClassFigi string       `db:"share_class_figi"`
	Symbol         string       `db:"symbol"`
	Symbol2        string       `db:"symbol2"`
	EquityType     string       `db:"equity_type"`
	CreatedAt      time.Time    `db:"created_at"`
	UpdatedAt      sql.NullTime `db:"updated_at"`
	DeletedAt      sql.NullTime `db:"deleted_at"`
}

func NewUsStockDataEntity(stockId uuid.UUID, currency string, description string,
	displaySymbol string, figi string, isIn string, mic string,
	shareClassFigi string, symbol string, symbol2 string, equityType string) *UsStock {
	return &UsStock{
		StockId:        stockId,
		Currency:       currency,
		Description:    description,
		DisplaySymbol:  displaySymbol,
		Figi:           figi,
		IsIn:           isIn,
		Mic:            mic,
		ShareClassFigi: shareClassFigi,
		Symbol:         symbol,
		Symbol2:        symbol2,
		EquityType:     equityType,
		CreatedAt:      time.Now(),
	}

}

func (u *UsStock) GetFigi() string {
	return u.Figi
}

func (u *UsStock) StockDataEntityToDomainEntity() *entity.UsStock {
	return &entity.UsStock{
		StockId:        u.StockId,
		Currency:       u.Currency,
		Description:    u.Description,
		DisplaySymbol:  u.DisplaySymbol,
		Figi:           u.Figi,
		IsIn:           u.IsIn,
		Mic:            u.Mic,
		ShareClassFigi: u.ShareClassFigi,
		Symbol:         u.Symbol,
		Symbol2:        u.Symbol2,
		EquityType:     u.EquityType,
	}
}
