package entity

import (
	"github.com/google/uuid"
	"time"
)

type UsStock struct {
	StockId        uuid.UUID `json:"stockId"`
	Currency       string    `json:"currency"`
	Description    string    `json:"description"`
	DisplaySymbol  string    `json:"displaySymbol"`
	Figi           string    `json:"figi"`
	IsIn           string    `json:"isIn"`
	ShareClassFigi string    `json:"shareClassFigi"`
	Symbol         string    `json:"symbol"`
	Symbol2        string    `json:"symbol2"`
	EquityType     string    `json:"equityType"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	DeletedAt      time.Time `json:"deletedAt"`
}

func NewUsStock(currency string, description string, displaySymbol string,
	figi string, isIn string, shareClassFigi string, symbol string, symbol2 string, equityType string) *UsStock {
	return &UsStock{
		Currency:       currency,
		Description:    description,
		DisplaySymbol:  displaySymbol,
		Figi:           figi,
		IsIn:           isIn,
		ShareClassFigi: shareClassFigi,
		Symbol:         symbol,
		Symbol2:        symbol2,
		EquityType:     equityType,
	}
}
