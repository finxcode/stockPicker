package metadata

import (
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
	getUsStockMetaDataPort         out.GetUsStockMetaDataPort
	checkStockExistInCache         cache.CheckStockExistRepository
	checkStockExistInDB            db.CheckStockExistRepository
	saveUsStockMetaDataInCache     cache.SaveUsStockMetaDataRepository
	SaveUsStockMetaDataInDB        db.SaveUsStockMetaDataRepository
	getUsStockByFigiInDbRepository db.GetUsStockByFigiRepository
}

func NewUsStockSymbolService(
	getUsStockMetaData out.GetUsStockMetaDataPort,
	checkStockExistInCache cache.CheckStockExistRepository,
	checkStockExistInDB db.CheckStockExistRepository,
	saveUsStockMetaDataInCache cache.SaveUsStockMetaDataRepository,
	SaveUsStockMetaDataInDB db.SaveUsStockMetaDataRepository,
	getUsStockByFigiInDbRepository db.GetUsStockByFigiRepository) *usStockMetaDataService {
	return &usStockMetaDataService{
		getUsStockMetaDataPort:         getUsStockMetaData,
		checkStockExistInCache:         checkStockExistInCache,
		checkStockExistInDB:            checkStockExistInDB,
		saveUsStockMetaDataInCache:     saveUsStockMetaDataInCache,
		SaveUsStockMetaDataInDB:        SaveUsStockMetaDataInDB,
		getUsStockByFigiInDbRepository: getUsStockByFigiInDbRepository,
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

func (u *usStockMetaDataService) saveToCache(stock *entity.UsStock) bool {
	return u.saveUsStockMetaDataInCache.SaveUsStockMetaDataInCache(stock)
}

func (u *usStockMetaDataService) saveToDB(stock *entity.UsStock) bool {
	return u.SaveUsStockMetaDataInDB.SaveUsStockMetaDataInDB(stock)
}

func (u *usStockMetaDataService) getStockInDb(figi string) (*entity.UsStock, error) {
	return u.getUsStockByFigiInDbRepository.GetUsStockByFigi(figi)
}

func (u *usStockMetaDataService) SaveUsStockMetaData() (int, int) {
	counterCache := 0
	counterDB := 0
	stocks := u.getUsStockMetaData()

	if stocks == nil || len(*stocks) == 0 {
		return 0, 0
	}

	for _, stock := range *stocks {
		if u.checkExistInCache(&stock) {
			continue
		} else {
			if u.checkExistInDB(&stock) {
				s, err := u.getStockInDb(stock.Figi)
				if err != nil {
					continue
				}
				if u.saveToCache(s) {
					counterCache++
				}
			} else {
				if u.saveToDB(&stock) {
					counterDB++
				}
				if u.saveToCache(&stock) {
					counterCache++
				}
			}
		}
	}
	return counterCache, counterDB
}
