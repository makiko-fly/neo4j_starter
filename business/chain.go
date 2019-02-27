package business

import (
	"strings"

	"gitlab.wallstcn.com/matrix/xgbkb/types"
)

var listChainsStmt = "MATCH (c:Chain) RETURN c SKIP $offset LIMIT $limit"

func ListChains(page, limit int64) (interface{}, error) {
	offset := (page - 1) * limit
	paramsMap := make(map[string]interface{})
	paramsMap["offset"] = offset
	paramsMap["limit"] = limit
	return Neo4jSingleQuery(listChainsStmt, paramsMap, false)
}

var getChainStmt = "MATCH (c:Chain {name:$name}) WHERE id(c) = $id RETURN c"

func GetChain(id int64, name string) (interface{}, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["id"] = id
	paramsMap["name"] = name
	return Neo4jSingleQuery(getChainStmt, paramsMap, false)
}

var createChainStmt = "CREATE (c:Chain {name: $name}) RETURN c"

func CreateChain(chainIn *types.ChainIn) (interface{}, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["name"] = chainIn.Name
	return Neo4jSingleQuery(createChainStmt, paramsMap, false)
}

func IsValidChainName(name string) bool {
	if strings.TrimSpace(name) == "" {
		return false
	}
	return true
}

var updateChainStmt = `
	MATCH (c:Chain {name: $oldName}) WHERE id(c) = $id
	SET c.name = $newName, c.previewImg = $previewImg, c.customData = $customData
	RETURN c
`

func UpdateChain(updateChainIn *types.UpdateChainIn) (interface{}, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["id"] = updateChainIn.Id
	paramsMap["oldName"] = updateChainIn.OldName
	paramsMap["newName"] = updateChainIn.Name
	paramsMap["previewImg"] = updateChainIn.PreviewImg
	paramsMap["customData"] = updateChainIn.CustomData
	return Neo4jSingleQuery(updateChainStmt, paramsMap, false)
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
	return Neo4jSingleQuery(addProductToChainStmt, paramsMap, false)
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
	return Neo4jSingleQuery(remProductToChainStmt, paramsMap, false)
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
	return Neo4jSingleQuery(getProductsOfChainStmt, paramsMap, true)
}
