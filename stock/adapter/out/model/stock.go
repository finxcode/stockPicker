package model

import (
	"github.com/google/uuid"
	"time"
)

type usStock struct {
	StockId        uuid.UUID `json:"stockId"`
	Currency       string    `json:"currency"`
	Description    string    `json:"description"`
	DisplaySymbol  string    `json:"displaySymbol"`
	Figi           string    `json:"figi"`
	IsIn           string    `json:"isIn"`
	Mic            string    `json:"mic"`
	ShareClassFigi string    `json:"shareClassFigi"`
	Symbol         string    `json:"symbol"`
	Symbol2        string    `json:"symbol2"`
	EquityType     string    `json:"equityType"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	DeletedAt      time.Time `json:"deletedAt"`
}

func NewUsStockDataEntity(stockId uuid.UUID, currency string, description string,
	displaySymbol string, figi string, isIn string, mic string,
	shareClassFigi string, symbol string, symbol2 string, equityType string) *usStock {
	return &usStock{
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
