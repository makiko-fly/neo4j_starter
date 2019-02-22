package business

import "gitlab.wallstcn.com/matrix/xgbkb/types"

var listChainsStmt = "MATCH (c:Chain) RETURN c SKIP $offset LIMIT $limit"

func ListChains(page, limit int64) (interface{}, error) {
	offset := (page - 1) * limit
	paramsMap := make(map[string]interface{})
	paramsMap["offset"] = offset
	paramsMap["limit"] = limit
	return QueryNeo4j(listChainsStmt, paramsMap, false)
}

var createChainStmt = "CREATE (c:Chain {name: $name}) RETURN c"

func CreateChain(chainIn *types.ChainIn) (interface{}, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["name"] = chainIn.Name
	return QueryNeo4j(createChainStmt, paramsMap, false)
}

func IsValidChainName(name string) bool {
	return true
}

var updateChainStmt = `
	MATCH (c:Chain {name: $oldName}) WHERE id(c) = $id
	SET c.name = $newName, c.previewImg = $previewImg
	RETURN c
`

func UpdateChain(updateChainIn *types.UpdateChainIn) (interface{}, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["id"] = updateChainIn.Id
	paramsMap["oldName"] = updateChainIn.OldName
	paramsMap["newName"] = updateChainIn.Name
	paramsMap["previewImg"] = updateChainIn.PreviewImg
	return QueryNeo4j(updateChainStmt, paramsMap, false)
}

var addProductToChainStmt = `
	MATCH (p:Product {name:$productName}), (c:Chain {name:$chainName}) 
	WHERE id(p) = $productId AND id(c) = $chainId 
	MERGE (p)-[r:WITHIN_CHAIN]->(c)
	RETURN p,r,c
`

func AddProductToChain(addProductToChainIn *types.AddProductToChainIn) (interface{}, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["productName"] = addProductToChainIn.ProductName
	paramsMap["productId"] = addProductToChainIn.ProductId
	paramsMap["chainName"] = addProductToChainIn.ChainName
	paramsMap["chainId"] = addProductToChainIn.ChainId
	return QueryNeo4j(addProductToChainStmt, paramsMap, false)
}

var remProductToChainStmt = `
	MATCH (p:Product {name:$productName}), (c:Chain {name:$chainName}) 
	WHERE id(p) = $productId AND id(c) = $chainId 
	MATCH (p)-[r:WITHIN_CHAIN]->(c)
	DELETE r
`

func RemProductFromChain(remProductFromChainIn *types.RemProductFromChainIn) (interface{}, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["productName"] = remProductFromChainIn.ProductName
	paramsMap["productId"] = remProductFromChainIn.ProductId
	paramsMap["chainName"] = remProductFromChainIn.ChainName
	paramsMap["chainId"] = remProductFromChainIn.ChainId
	return QueryNeo4j(remProductToChainStmt, paramsMap, false)
}

var getProductsOfChainStmt = `
	MATCH (c:Chain) WHERE id(c) = $chainId AND c.name = $chainName
	MATCH (c)<-[:WITHIN_CHAIN]-(p)
	RETURN p
`

func GetProductsOfChain(chainId int64, chainName string) (interface{}, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["chainId"] = chainId
	paramsMap["chainName"] = chainName
	return QueryNeo4j(getProductsOfChainStmt, paramsMap, true)
}
