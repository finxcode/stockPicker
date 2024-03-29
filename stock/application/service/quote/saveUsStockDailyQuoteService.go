package quote

import (
	"fmt"
	"github.com/google/uuid"
	"stockPicker/stock/application/port/out"
	"stockPicker/stock/application/port/out/cache"
	"stockPicker/stock/application/port/out/db"
	"stockPicker/stock/init/config"
	"strings"
	"time"
)

type saveUsStockDailyQuoteService struct {
	config                          *config.Config
	getUsStockDailyQuotePort        out.GetUsStockDailyQuotePort
	getAllUsStockFigi               cache.GetAllUsStockFigiRepository
	getUsStocksByFigi               cache.GetUsStocksByFigiRepository
	saveUsStockDailyQuoteRepository db.SaveUsStockDailyQuoteRepository
}

func NewSaveUsStockDailyQuoteService(
	config *config.Config,
	getUsStockDailyQuotePort out.GetUsStockDailyQuotePort,
	getAllUsStockFigi cache.GetAllUsStockFigiRepository,
	getUsStocksByFigi cache.GetUsStocksByFigiRepository,
	saveUsStockDailyQuoteRepository db.SaveUsStockDailyQuoteRepository) *saveUsStockDailyQuoteService {
	return &saveUsStockDailyQuoteService{
		config:                          config,
		getUsStockDailyQuotePort:        getUsStockDailyQuotePort,
		getAllUsStockFigi:               getAllUsStockFigi,
		getUsStocksByFigi:               getUsStocksByFigi,
		saveUsStockDailyQuoteRepository: saveUsStockDailyQuoteRepository,
	}
}

func (s *saveUsStockDailyQuoteService) SaveUsStockDailyQuotes() int {
	countStockDailyQuote := 0
	figis := s.getAllUsStockFigi.GetAllUsStockFigi()
	if figis == nil {
		return 0
	}
	for _, figi := range *figis {
		res := s.getUsStocksByFigi.GetUsStockSymbolByFigi(figi)
		if res == "" {
			continue
		}
		quote, err := s.getUsStockDailyQuotePort.GetUsStockDailyQuote(urlQuoteBuilder(
			s.config.Xueqiu.BaseUrl, getSymbolAndStockId(res)[0]), getSymbolAndStockId(res)[0])
		if err != nil {
			time.Sleep(time.Second * 3)
			continue
		}
		quote.SetStockId(uuid.MustParse(getSymbolAndStockId(res)[1]))

		err = s.saveUsStockDailyQuoteRepository.SaveUsStockDailyQuoteToDb(quote)
		if err != nil {
			continue
		}
		countStockDailyQuote++

		time.Sleep(time.Second * 3)
	}

	return countStockDailyQuote
}

func urlQuoteBuilder(url, symbol string) string {
	return fmt.Sprintf("%s/%s", url, symbol)
}

func getSymbolAndStockId(res string) []string {
	return strings.Split(res, " ")
}
