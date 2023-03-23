package mapper

import (
	"github.com/google/uuid"
	"time"
)

type StockDailyPrice struct {
	ID                 uint      `json:"ID"`
	StockId            uuid.UUID `json:"stockId"`
	Symbol             string    `json:"symbol"`             //symbol
	Current            float32   `json:"current"`            //current
	LastClose          float32   `json:"lastClose"`          //last_close
	High               float32   `json:"high"`               //high
	Low                float32   `json:"low"`                //low
	Open               float32   `json:"open"`               //open
	Change             float32   `json:"change"`             //chg
	Percent            float32   `json:"percent"`            //percent
	High52Week         float32   `json:"high52Week"`         //high52w
	Low52Week          float32   `json:"low52Week"`          //low52w
	Volume             uint      `json:"volume"`             //volume
	Amount             float32   `json:"amount"`             //amount
	MarketCapital      uint      `json:"marketCapital"`      //market_capital
	TotalShares        uint      `json:"totalShares"`        //total_shares
	Dividend           float32   `json:"dividend"`           //dividend
	DividendYield      float32   `json:"dividendYield"`      //dividend_yield
	EPS                float32   `json:"EPS"`                //eps
	TurnoverRate       float32   `json:"turnoverRate"`       //turnover_rate
	VolumeRatio        float32   `json:"volumeRatio"`        //volume_ratio
	Amplitude          float32   `json:"amplitude"`          //amplitude
	CurrentYearPercent float32   `json:"currentYearPercent"` //current_year_percent
	IssueDate          time.Time `json:"issueDate"`          //issue_date
	PeTtm              float32   `json:"peTtm"`              //pe_ttm
	PeLyr              float32   `json:"peLyr"`              //pe_lyr
	PeForecast         float32   `json:"peForecast"`         //pe_forecast
	Navps              float32   `json:"navps"`              //navps
	Pb                 float32   `json:"pb"`                 //pb
	Psr                float32   `json:"psr"`                //psr
	TradingDay         string    `json:"tradingDay"`         //timestamp
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	DeletedAt          time.Time `json:"deletedAt"`
}
