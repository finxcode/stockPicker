package fetcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	finnHubConfig "stockPicker/ext/finnhub/config"
	"stockPicker/ext/finnhub/entity"
)

var ErrRequestError error
var ErrParseJsonError error

type StockSymbolFetcher struct {
	config *finnHubConfig.Config
}

func NewStockSymbolFetcher() *StockSymbolFetcher {
	if err, c := finnHubConfig.New(); err != nil {
		return nil
	} else {
		return &StockSymbolFetcher{
			config: c,
		}
	}
}

func (s *StockSymbolFetcher) GetUsStockSymbol() (*[]entity.Symbol, error) {
	var symbols []entity.Symbol

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, buildStockSymbolUrl(s.config), nil)
	if err != nil {
		ErrRequestError = errors.New("http request error")
		return nil, ErrRequestError
	}

	q := req.URL.Query()
	q.Add("exchange", "US")
	q.Add("token", s.config.ApiKey)

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		ErrRequestError = errors.New("client.Do error")
		return nil, ErrRequestError
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ErrRequestError = errors.New("read response body error")
		return nil, ErrRequestError
	}
	err = json.Unmarshal(responseBody, &symbols)
	if err != nil {
		ErrParseJsonError = fmt.Errorf("parse json failed with error: %s", err.Error())
		return nil, ErrParseJsonError
	}
	return &symbols, nil
}

func buildStockSymbolUrl(c *finnHubConfig.Config) string {
	return fmt.Sprintf("%s%s%s",
		c.FinnHub.BaseUrl,
		c.FinnHub.Stock,
		c.FinnHub.Symbol)
}
