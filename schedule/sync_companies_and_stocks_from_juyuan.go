package schedule

import (
	"time"

	"gitlab.wallstcn.com/baoer/matrix/xgbkb/business"
	"gitlab.wallstcn.com/baoer/matrix/xgbkb/g"
	"gitlab.wallstcn.com/baoer/matrix/xgbkb/std/redislogger"
	"gitlab.wallstcn.com/baoer/matrix/xgbkb/types"
)

type CompanyAndStockInfo struct {
	SecuCode    string `gorm:"column:SecuCode"`
	SecuAbbr    string `gorm:"column:SecuAbbr"`
	SecuMarket  int64  `gorm:"column:SecuMarket"`
	CompanyCode int64  `gorm:"column:CompanyCode"`
	ChiName     string `gorm:"column:ChiName"`
	ChiNameAbbr string `gorm:"column:ChiNameAbbr"`
}

func SyncCompaniesAndStocksFromJuyuan() error {
	redislogger.Printf("=============== SyncCompaniesAndStocksFromJuyuan Started At: %v =================", time.Now())
	// get all companies and stocks from juyuan db
	var companyAndStockInfoRecArr []*CompanyAndStockInfo
	if err := g.JuyuanDb.Table("SecuMain").Select("SecuCode, SecuAbbr, SecuMarket, CompanyCode, ChiName, ChiNameAbbr").
		Where("SecuMarket in (83, 90) AND SecuCategory = 1 AND ListedState = 1").
		Find(&companyAndStockInfoRecArr).Error; err != nil {
		redislogger.Errf("SyncCompaniesAndStocksFromJuyuan, fails to get company and stock info from db, err: %v", err)
		return err
	}
	// sync companies and stocks
	for _, companyAndStockInfoRec := range companyAndStockInfoRecArr {
		companyIn := &types.CompanyIn{
			Name:     companyAndStockInfoRec.ChiName,
			NameAbbr: companyAndStockInfoRec.ChiNameAbbr,
			Code:     companyAndStockInfoRec.CompanyCode,
		}
		if _, err := business.CreateCompany(companyIn); err != nil {
			redislogger.Errf("SyncCompaniesAndStocksFromJuyuan, fails to create company, err: %v", err)
		}
		stockIn := &types.StockIn{
			Symbol: business.StockSymFromSecuCodeAndMarket(companyAndStockInfoRec.SecuCode,
				companyAndStockInfoRec.SecuMarket),
			Name: companyAndStockInfoRec.SecuAbbr,
		}
		if _, err := business.MergeStock(stockIn); err != nil {
			redislogger.Errf("SyncCompaniesAndStocksFromJuyuan, fails to merge stock, err: %v", err)
		}
		// create relationship bewteen them
		if _, err := business.MergeListedAsRelation(companyIn, stockIn); err != nil {
			redislogger.Errf("SyncCompaniesAndStocksFromJuyuan, fails to create ListedAs relation, err: %v", err)
		}
	}

	redislogger.Printf("=============== SyncCompaniesAndStocksFromJuyuan Finished At: %v ===============", time.Now())
	return nil
}
