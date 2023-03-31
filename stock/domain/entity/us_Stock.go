package entity

import (
	"github.com/google/uuid"
)

type UsStock struct {
	StockId        uuid.UUID
	Currency       string
	Description    string
	DisplaySymbol  string
	Figi           string
	IsIn           string
	Mic            string
	ShareClassFigi string
	Symbol         string
	Symbol2        string
	EquityType     string
}

func NewUsStock(currency string, description string, displaySymbol string,
	figi string, isIn string, mic string, shareClassFigi string,
	symbol string, symbol2 string, equityType string) *UsStock {
	return &UsStock{
		StockId:        uuid.New(),
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
	}
}

func (u *UsStock) SetId() {
	u.StockId = uuid.New()
}
