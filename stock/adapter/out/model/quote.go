package model

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type UsStockDailyQuote struct {
	ID                 uint         `json:"ID" db:"id"`
	StockId            uuid.UUID    `json:"stock_id" db:"stock_id"`
	Symbol             string       `json:"symbol" db:"symbol"`                                                                 //symbol
	Current            float32      `json:"current" mapstructure:"current" db:"current"`                                        //current
	LastClose          float32      `json:"last_close" mapstructure:"last_close" db:"last_close"`                               //last_close
	High               float32      `json:"high" mapstructure:"high" db:"high"`                                                 //high
	Low                float32      `json:"low" mapstructure:"low" db:"low"`                                                    //low
	Open               float32      `json:"open" mapstructure:"open" db:"open"`                                                 //open
	Change             float32      `json:"change" mapstructure:"chg" db:"chg"`                                                 //chg
	Percent            float32      `json:"percent" mapstructure:"percent" db:"percent"`                                        //percent
	High52Week         float32      `json:"high52w" mapstructure:"high52w" db:"high52w"`                                        //high52w
	Low52Week          float32      `json:"low52w" mapstructure:"low52w" db:"low52w"`                                           //low52w
	Volume             uint         `json:"volume" mapstructure:"volume" db:"volume"`                                           //volume
	Amount             float32      `json:"amount" mapstructure:"amount" db:"amount"`                                           //amount
	MarketCapital      uint         `json:"market_capital" mapstructure:"market_capital" db:"market_capital"`                   //market_capital
	TotalShares        uint         `json:"total_shares" mapstructure:"total_shares" db:"total_shares"`                         //total_shares
	Dividend           float32      `json:"dividend" mapstructure:"dividend" db:"dividend"`                                     //dividend
	DividendYield      float32      `json:"dividend_yield" mapstructure:"dividend_yield" db:"dividend_yield"`                   //dividend_yield
	EPS                float32      `json:"EPS" mapstructure:"eps" db:"eps"`                                                    //eps
	TurnoverRate       float32      `json:"turnover_rate" mapstructure:"turnover_rate" db:"turnover_rate"`                      //turnover_rate
	VolumeRatio        float32      `json:"volume_ratio" mapstructure:"volume_ratio" db:"volume_ratio"`                         //volume_ratio
	Amplitude          float32      `json:"amplitude" mapstructure:"amplitude" db:"amplitude"`                                  //amplitude
	CurrentYearPercent float32      `json:"current_year_percent" mapstructure:"current_year_percent" db:"current_year_percent"` //current_year_percent
	IssueDate          int          `json:"issue_date" mapstructure:"issue_date" db:"issue_date"`                               //issue_date
	PeTtm              float32      `json:"pe_ttm" mapstructure:"pe_ttm" db:"pe_ttm"`                                           //pe_ttm
	PeLyr              float32      `json:"pe_lyr" mapstructure:"pe_lyr" db:"pe_lyr"`                                           //pe_lyr
	PeForecast         float32      `json:"pe_forecast" mapstructure:"pe_forecast" db:"pe_forecast"`                            //pe_forecast
	Navps              float32      `json:"navps" mapstructure:"navps" db:"navps"`                                              //navps
	Pb                 float32      `json:"pb" mapstructure:"pb" db:"pb"`                                                       //pb
	Psr                float32      `json:"psr" mapstructure:"psr" db:"psr"`                                                    //psr
	FetchTime          int          `json:"timestamp" mapstructure:"timestamp" db:"fetch_time"`                                 //timestamp
	TradingDay         string       `json:"trading_day" db:"trading_day"`
	CreatedAt          time.Time    `json:"createdAt" db:"created_at"`
	UpdatedAt          sql.NullTime `json:"updatedAt" db:"updated_at"`
	DeletedAt          sql.NullTime `json:"deletedAt" db:"deleted_at"`
}

func NewStockDailyQuote(symbol string, current float32, lastClose float32,
	high float32, low float32, open float32, change float32, percent float32, high52Week float32,
	low52Week float32, volume uint, amount float32, marketCapital uint, totalShares uint,
	dividend float32, dividendYield float32, EPS float32, turnoverRate float32, volumeRatio float32,
	amplitude float32, currentYearPercent float32, issueDate int, peTtm float32, peLyr float32,
	peForecast float32, navps float32, pb float32, psr float32, timestamp int, tradingDay string) *UsStockDailyQuote {
	return &UsStockDailyQuote{Symbol: symbol, Current: current, LastClose: lastClose,
		High: high, Low: low, Open: open, Change: change, Percent: percent, High52Week: high52Week, Low52Week: low52Week,
		Volume: volume, Amount: amount, MarketCapital: marketCapital, TotalShares: totalShares, Dividend: dividend,
		DividendYield: dividendYield, EPS: EPS, TurnoverRate: turnoverRate, VolumeRatio: volumeRatio, Amplitude: amplitude,
		CurrentYearPercent: currentYearPercent, IssueDate: issueDate, PeTtm: peTtm, PeLyr: peLyr, PeForecast: peForecast,
		Navps: navps, Pb: pb, Psr: psr, FetchTime: timestamp, TradingDay: tradingDay, CreatedAt: time.Now()}
}

func (s *UsStockDailyQuote) SetTradingDay(timestamp int64) {
	var cstSh, _ = time.LoadLocation("EST")
	dateTime := time.Unix(0, timestamp*int64(time.Millisecond)).In(cstSh).Format("2006-01-02 15:04:05")
	s.TradingDay = dateTime[0:10]
}

func (s *UsStockDailyQuote) SetStockId(stockId uuid.UUID) {
	s.StockId = stockId
}

func (s *UsStockDailyQuote) SetSymbol(symbol string) {
	s.Symbol = symbol
}
