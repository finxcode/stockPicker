package parser

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	xueQiuConfig "stockPicker/ext/xueqiu/config"
	"strconv"
	"strings"
)

type Token struct {
	Name string
	Type string
}

type StockQuoteParser struct {
	config *xueQiuConfig.Config
}

func NewStockPriceParser() *StockQuoteParser {
	c, err := xueQiuConfig.New()
	if err != nil {
		return nil
	} else {
		return &StockQuoteParser{
			config: c,
		}
	}
}

func buildTokens() *[28]Token {
	tokens := [28]Token{
		{"current", "floatQuote"},
		{"last_close", "float"},
		{"high", "float"},
		{"low", "float"},
		{"open", "float"},
		{"chg", "floatQuote"},
		{"percent", "float"},
		{"high52w", "float"},
		{"low52w", "float"},
		{"volume", "uint"},
		{"amount", "float"},
		{"market_capital", "uint"},
		{"total_shares", "uint"},
		{"dividend", "float"},
		{"dividend_yield", "float"},
		{"eps", "float"},
		{"turnover_rate", "float"},
		{"volume_ratio", "float"},
		{"amplitude", "float"},
		{"current_year_percent", "float"},
		{"issue_date", "int"},
		{"pe_ttm", "float"},
		{"pe_lyr", "float"},
		{"pe_forecast", "float"},
		{"navps", "float"},
		{"pb", "float"},
		{"psr", "float"},
		{"timestamp", "int"},
	}
	return &tokens
}

func parseToFloat(text *string, tokenName string, sep string) float32 {
	pattern := &regexp.Regexp{}
	if sep == "" {
		pattern = regexp.MustCompile(
			fmt.Sprintf("\"%s\":[-+]?[0-9]+.?[0-9]*", tokenName))
	} else {
		pattern = regexp.MustCompile(fmt.Sprintf("\"%s\":%s[-+]?[0-9]+.?[0-9]*%s", tokenName, sep, sep))
	}
	strWithTokenName := pattern.FindString(*text)
	if len(strWithTokenName) < len(tokenName) {
		return 0
	}
	patternNumber := regexp.MustCompile("[-+]?[0-9]+.?[0-9]+|[+-]?[0-9]+")
	content := patternNumber.FindString(strWithTokenName)
	if len(content) == 0 {
		return 0
	}
	number, err := strconv.ParseFloat(content, 32)

	if err != nil {
		log.Println(fmt.Sprintf("error occurred when parse token %s: %s", tokenName, err.Error()))
	}
	return float32(number)
}

func parseToInt(text *string, tokenName string) int {
	pattern := regexp.MustCompile(fmt.Sprintf("\"%s\":[+-]?[0-9]+", tokenName))
	strWithTokenName := pattern.FindString(*text)
	patternNumber := regexp.MustCompile("[+-]?[0-9]+")
	content := patternNumber.FindString(strWithTokenName)
	if len(content) == 0 {
		return 0
	}
	number, err := strconv.Atoi(content)
	if err != nil {
		log.Println(fmt.Sprintf("error occurred when parse token %s: %s", tokenName, err.Error()))
	}
	return number
}

func getXueQiuStockPageHtml(url string) (*string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	html := string(b)
	return &html, nil
}

func (p *StockQuoteParser) ParseStockData(url string) (map[string]interface{}, error) {
	html, err := getXueQiuStockPageHtml(url)
	if err != nil {
		return make(map[string]interface{}), err
	}

	res := make(map[string]interface{})
	tokens := buildTokens()

	startIndex := strings.Index(*html, "quote: {")
	endIndex := strings.Index(*html, "quoteTags: [{")
	if startIndex >= endIndex {
		return make(map[string]interface{}), errors.New("could not find target text in html body")
	}

	text := (*html)[startIndex:endIndex]
	for _, token := range *tokens {
		if token.Type == "floatQuote" {
			res[token.Name] = parseToFloat(&text, token.Name, "\"")
		} else if token.Type == "float" {
			res[token.Name] = parseToFloat(&text, token.Name, "")
		} else if token.Type == "uint" || token.Type == "int" {
			res[token.Name] = parseToInt(&text, token.Name)
		}
	}
	return res, nil
}
