package out

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	dataEntity "stockPicker/stock/adapter/out/entity"
	"stockPicker/stock/domain/entity"
)

type saveUsStockMetaDataConsoleController struct {
	db *sqlx.DB
}

func NewSaveUsStockMetaDataConsoleController() *saveUsStockMetaDataConsoleController {
	return &saveUsStockMetaDataConsoleController{}
}

func (s *saveUsStockMetaDataConsoleController) SaveUsStockMetaDataInCache(stock *entity.UsStock) bool {
	return false
}

func (s *saveUsStockMetaDataConsoleController) SaveUsStockMetaDataInDB(domainStock *entity.UsStock) bool {
	stock := dataEntity.NewUsStockDataEntity(uuid.New(), domainStock.Currency,
		domainStock.Description, domainStock.DisplaySymbol,
		domainStock.Figi, domainStock.IsIn, domainStock.Mic,
		domainStock.ShareClassFigi, domainStock.Symbol, domainStock.Symbol2, domainStock.EquityType)
	query := ""
	return false
}
