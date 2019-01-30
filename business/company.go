package business

import "gitlab.wallstcn.com/matrix/xgbkb/types"

var createCompanyStmt = "CREATE (c:Company {name: $name, nameAbbr: $nameAbbr, code: $code}) RETURN c"

func CreateCompany(companyIn *types.CompanyIn) (interface{}, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["name"] = companyIn.Name
	paramsMap["nameAbbr"] = companyIn.NameAbbr
	paramsMap["code"] = companyIn.Code
	return QueryNeo4j(createCompanyStmt, paramsMap, false)
}

var mergeListedAsStmt = `
	MATCH (c:Company {code: $code}), (s:Stock {symbol: $symbol})
	MERGE (c)-[r:LISTED_AS]->(s)
	RETURN r
`

func MergeListedAsRelation(companyIn *types.CompanyIn, stockIn *types.StockIn) (interface{}, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["code"] = companyIn.Code
	paramsMap["symbol"] = stockIn.Symbol
	return QueryNeo4j(mergeListedAsStmt, paramsMap, false)
}
