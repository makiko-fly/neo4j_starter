package business

var getDirectlyRelatedNodesStmt = "MATCH (a)-[]-(b) WHERE id(a) = $nodeId RETURN b, labels(b)"

func GetDirectlyRelatedNodes(nodeId int64) (interface{}, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["nodeId"] = nodeId
	return QueryNeo4j(getDirectlyRelatedNodesStmt, paramsMap, false)
}
