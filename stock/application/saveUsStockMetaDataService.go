package application

import (
	"fmt"
	"stockPicker/config"
	"stockPicker/stock/application/port/out"
	"stockPicker/stock/application/port/out/cache"
	"stockPicker/stock/application/port/out/db"
	"stockPicker/stock/domain/entity"
)

const (
	XNAS = "XNAS"
	XNYS = "XNYS"
	XASE = "XASE"
	Type = "Common Stock"
)

type usStockMetaDataService struct {
	config                     *config.Config
	getUsStockMetaDataPort     out.GetUsStockMetaDataPort
	checkStockExistInCache     cache.CheckStockExistRepository
	checkStockExistInDB        db.CheckStockExistRepository
	saveUsStockMetaDataInCache cache.SaveUsStockMetaDataRepository
	SaveUsStockMetaDataInDB    db.SaveUsStockMetaDataRepository
}

func NewUsStockSymbolService(config *config.Config,
	getUsStockMetaData out.GetUsStockMetaDataPort,
	checkStockExistInCache cache.CheckStockExistRepository,
	checkStockExistInDB db.CheckStockExistRepository,
	saveUsStockMetaDataInCache cache.SaveUsStockMetaDataRepository,
	SaveUsStockMetaDataInDB db.SaveUsStockMetaDataRepository) *usStockMetaDataService {
	return &usStockMetaDataService{
		config:                     config,
		getUsStockMetaDataPort:     getUsStockMetaData,
		checkStockExistInCache:     checkStockExistInCache,
		checkStockExistInDB:        checkStockExistInDB,
		saveUsStockMetaDataInCache: saveUsStockMetaDataInCache,
		SaveUsStockMetaDataInDB:    SaveUsStockMetaDataInDB,
	}
}

// filter criteria
// mic + type
// 1. XNAS + Common Stock
// 2. XNYS + Common Stock
// 3. XASE + Common Stock

func (u *usStockMetaDataService) getUsStockMetaData() *[]entity.UsStock {
	var usCommonStock []entity.UsStock
	stocks := u.getUsStockMetaDataPort.GetUsStockMetaData()
	fmt.Println(len(*stocks))
	for _, stock := range *stocks {
		if (stock.Mic == XNYS || stock.Mic == XNAS || stock.Mic == XASE) && stock.EquityType == Type {
			usCommonStock = append(usCommonStock, stock)
		}
	}
	return &usCommonStock
}

func (u *usStockMetaDataService) checkExistInCache(stock *entity.UsStock) bool {
	return u.checkStockExistInCache.CheckStockExistInCache(stock.Figi)
}

func (u *usStockMetaDataService) checkExistInDB(stock *entity.UsStock) bool {
	return u.checkStockExistInDB.CheckStockExistInDB(stock.Figi)
}

func (u *usStockMetaDataService) SaveUsStockMetaData(stocks *[]entity.UsStock) (int, int) {
	counterCache := 0
	counterDB := 0

	for _, stock := range *stocks {
		if u.checkExistInCache(&stock) {
			continue
		} else {
			if u.saveUsStockMetaDataInCache.SaveUsStockMetaDataInCache(&stock) {
				counterCache++
			}
			if u.checkExistInDB(&stock) {
				continue
			} else {
				if u.SaveUsStockMetaDataInDB.SaveUsStockMetaDataInDB(&stock) {
					counterDB++
				}
			}
		}
	}
	return counterCache, counterDB
}
