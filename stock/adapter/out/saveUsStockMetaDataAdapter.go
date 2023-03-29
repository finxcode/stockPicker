package out

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"stockPicker/stock/domain/entity"
)

type saveUsStockMetaDataConsoleController struct {
	rds *redis.Client
	db  *sqlx.DB
}

func NewSaveUsStockMetaDataConsoleController(rds *redis.Client, db *sqlx.DB) *saveUsStockMetaDataConsoleController {
	return &saveUsStockMetaDataConsoleController{
		rds: rds,
		db:  db,
	}
}

func (s *saveUsStockMetaDataConsoleController) SaveUsStockMetaDataInCache(stock *entity.UsStock) bool {

	return false
}

func (s *saveUsStockMetaDataConsoleController) SaveUsStockMetaDataInDB(domainStock *entity.UsStock) bool {
	//stock := dataEntity.NewUsStockDataEntity(uuid.New(), domainStock.Currency,
	//	domainStock.Description, domainStock.DisplaySymbol,
	//	domainStock.Figi, domainStock.IsIn, domainStock.Mic,
	//	domainStock.ShareClassFigi, domainStock.Symbol, domainStock.Symbol2, domainStock.EquityType)
	//query := "INSERT INTO us-stock-meta-data "
	return false
}
