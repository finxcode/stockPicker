package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"stockPicker/config"
	"stockPicker/global"
	"stockPicker/internal/domain/stock"
	"stockPicker/internal/repository"
)

var ErrRequestError error
var ErrParseJsonError error

type FinnHubFetcherRepository struct {
	config *config.Config
}

func NewFinnHubFetcherRepository(c *config.Config) repository.FetcherRepository {
	return &FinnHubFetcherRepository{
		config: c,
	}
}

func (f FinnHubFetcherRepository) FetchStockSymbol() (*[]stock.Symbol, error) {
	var symbols []stock.Symbol

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, buildStockSymbolUrl(f.config), nil)
	if err != nil {
		global.App.Logger.Error("build stock symbol request error", zap.String("http.NewRequest error", err.Error()))
		ErrRequestError = errors.New("http request error")
		return nil, ErrRequestError
	}

	q := req.URL.Query()
	q.Add("exchange", "US")
	q.Add("token", f.config.FinnHub.ApiKey)

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

func buildStockSymbolUrl(c *config.Config) string {
	return fmt.Sprintf("%s%s%s",
		c.FinnHub.BaseUrl,
		c.FinnHub.Stock,
		c.FinnHub.Symbol)
}
