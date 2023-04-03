package entity

import (
	"github.com/google/uuid"
)

type StockDailyQuote struct {
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
	IssueDate          int       `json:"issueDate"`          //issue_date
	PeTtm              float32   `json:"peTtm"`              //pe_ttm
	PeLyr              float32   `json:"peLyr"`              //pe_lyr
	PeForecast         float32   `json:"peForecast"`         //pe_forecast
	Navps              float32   `json:"navps"`              //navps
	Pb                 float32   `json:"pb"`                 //pb
	Psr                float32   `json:"psr"`                //psr
	Timestamp          int       `json:"timestamp"`          //timestamp
	TradingDay         string    `json:"tradingDay"`
}

func NewStockDailyQuote(symbol string, current float32, lastClose float32,
	high float32, low float32, open float32, change float32, percent float32, high52Week float32,
	low52Week float32, volume uint, amount float32, marketCapital uint, totalShares uint,
	dividend float32, dividendYield float32, EPS float32, turnoverRate float32, volumeRatio float32,
	amplitude float32, currentYearPercent float32, issueDate int, peTtm float32, peLyr float32,
	peForecast float32, navps float32, pb float32, psr float32, timestamp int, tradingDay string) *StockDailyQuote {
	return &StockDailyQuote{Symbol: symbol, Current: current, LastClose: lastClose,
		High: high, Low: low, Open: open, Change: change, Percent: percent, High52Week: high52Week, Low52Week: low52Week,
		Volume: volume, Amount: amount, MarketCapital: marketCapital, TotalShares: totalShares, Dividend: dividend,
		DividendYield: dividendYield, EPS: EPS, TurnoverRate: turnoverRate, VolumeRatio: volumeRatio, Amplitude: amplitude,
		CurrentYearPercent: currentYearPercent, IssueDate: issueDate, PeTtm: peTtm, PeLyr: peLyr, PeForecast: peForecast,
		Navps: navps, Pb: pb, Psr: psr, Timestamp: timestamp, TradingDay: tradingDay}
}

func (s *StockDailyQuote) SetStockId(stockId uuid.UUID) {
	s.StockId = stockId
}
