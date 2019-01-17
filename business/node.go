package business

import (
	"strings"
)

var getDirectlyRelatedNodesStmt = "MATCH (a)-[r]-(b) WHERE id(a) = $nodeId RETURN r"

func GetDirectlyRelatedNodes(nodeId int64) (interface{}, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["nodeId"] = nodeId
	return QueryNeo4j(getDirectlyRelatedNodesStmt, paramsMap, true)
}

// name's emptiness should alreay have been checked
// TODO.. add more rules
func IsValidNodeName(name string) bool {
	if strings.Contains(name, "\"") {
		return false
	} else if strings.Contains(name, "'") {
		return false
	}
	return true
}
