package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"stockPicker/global"
)

type StockSymbol struct {
	Currency       string `json:"currency"`
	Description    string `json:"description"`
	DisplaySymbol  string `json:"displaySymbol"`
	Figi           string `json:"figi"`
	Isin           string `json:"isin,omitempty"`
	ShareClassFigi string `json:"shareClassFIGI"`
	Symbol         string `json:"symbol"`
	Symbol2        string `json:"symbol2"`
	Type           string `json:"type"`
}

var ErrRequestError error
var ErrParseJsonError error

func FetchStockSymbol() (*[]StockSymbol, error) {
	var symbols []StockSymbol

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, buildStockSymbolUrl(), nil)
	if err != nil {
		global.App.Logger.Error("build stock symbol request error", zap.String("http.NewRequest error", err.Error()))
		ErrRequestError = errors.New("http request error")
		return nil, ErrRequestError
	}

	q := req.URL.Query()
	q.Add("exchange", "US")
	q.Add("token", global.App.Config.FinnHub.ApiKey)

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		global.App.Logger.Error("stock symbol request error", zap.String("clinet.Do error", err.Error()))
		ErrRequestError = errors.New("client.Do error")
		return nil, ErrRequestError
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.App.Logger.Error("read response body error", zap.String("read response body error", err.Error()))
		ErrRequestError = errors.New("read response body error")
		return nil, ErrRequestError
	}
	err = json.Unmarshal(responseBody, &symbols)
	if err != nil {
		global.App.Logger.Error("parse json error", zap.String("parse json error", err.Error()))
		ErrParseJsonError = fmt.Errorf("parse json failed with error: %s", err.Error())
		return nil, ErrParseJsonError
	}
	return &symbols, nil
}

func buildStockSymbolUrl() string {
	return fmt.Sprintf("%s%s%s",
		global.App.Config.BaseUrl,
		global.App.Config.FinnHub.Stock,
		global.App.Config.FinnHub.Symbol)
}
