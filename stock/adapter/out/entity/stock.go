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
	Mic            string    `json:"mic"`
	ShareClassFigi string    `json:"shareClassFigi"`
	Symbol         string    `json:"symbol"`
	Symbol2        string    `json:"symbol2"`
	EquityType     string    `json:"equityType"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	DeletedAt      time.Time `json:"deletedAt"`
}
