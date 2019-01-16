package business

import "gitlab.wallstcn.com/baoer/matrix/xgbkb/types"

var mergeStockStmt = `
	MERGE (s:Stock { symbol: $symbol })
	ON MATCH SET s.name = $name
	ON CREATE SET s.name = $name
	RETURN s
`

func MergeStock(stockIn *types.StockIn) (interface{}, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["symbol"] = stockIn.Symbol
	paramsMap["name"] = stockIn.Name
	return QueryNeo4j(mergeStockStmt, paramsMap, false)
}

// func CreateStock(stockIn *types.StockIn) (interface{}, error) {
// 	paramsMap := make(map[string]interface{})
// 	paramsMap["name"] = companyIn.Name
// 	paramsMap["nameAbbr"] = companyIn.NameAbbr
// 	paramsMap["code"] = companyIn.Code
// 	return QueryNeo4j(createCompanyStmt, paramsMap, false)
// }

// func UpdateStock(stockIn *types.Stock) (interface{}, error) {

// }

// var getStockBySymStmt = "MATCH (s:Stock) WHERE s.symbol = $sym RETURN s"

// func GetStockBySym(sym string) (*types.Stock, error) {

// }

func StockSymFromSecuCodeAndMarket(secuCode string, secuMarket int64) string {
	var suffix = ".SS"
	if secuMarket == 90 {
		suffix = ".SZ"
	}
	return secuCode + suffix
}
