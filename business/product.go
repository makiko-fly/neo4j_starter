package business

import (
	"gitlab.wallstcn.com/baoer/matrix/xgbkb/types"
)

var createProductStmt = "CREATE (p:Product {name: $name, imgActivated: $imgActivated, imgNormal: $imgNormal}) RETURN p"

func CreateProduct(productIn *types.ProductIn) (interface{}, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["name"] = productIn.Name
	paramsMap["imgActivated"] = productIn.ImgActivated
	paramsMap["imgNormal"] = productIn.ImgNormal
	return QueryNeo4j(createProductStmt, paramsMap, false)
}

var updateProductStmt = `
	MATCH (p:Product {name: $oldName}) WHERE id(p) = $id
	SET p.name = $newName, p.imgActivated = $imgActivated, p.imgNormal = $imgNormal
	RETURN p
`

func UpdateProduct(id int64, oldName string, productIn *types.ProductIn) (interface{}, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["id"] = id
	paramsMap["oldName"] = oldName
	paramsMap["newName"] = productIn.Name
	paramsMap["imgActivated"] = productIn.ImgActivated
	paramsMap["imgNormal"] = productIn.ImgNormal
	return QueryNeo4j(updateProductStmt, paramsMap, false)
}

func IsValidProductName(name string) bool {
	return true
}
