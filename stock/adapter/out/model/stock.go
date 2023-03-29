package model

import (
	"github.com/google/uuid"
	"time"
)

type UsStock struct {
	StockId        uuid.UUID `db:"stock_id"`
	Currency       string    `db:"currency"`
	Description    string    `db:"description"`
	DisplaySymbol  string    `db:"display_symbol"`
	Figi           string    `db:"figi"`
	IsIn           string    `db:"is_in"`
	Mic            string    `db:"mic"`
	ShareClassFigi string    `db:"share_class_figi"`
	Symbol         string    `db:"symbol"`
	Symbol2        string    `db:"symbol2"`
	EquityType     string    `db:"equity_type"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	DeletedAt      time.Time `db:"deleted_at"`
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
		EquityType:     equityType}
}
