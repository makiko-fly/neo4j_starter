package business

import (
	"fmt"
	"strings"
)

var getDirectlyRelatedNodesStmt = "MATCH (a {name:$nodeName})-[r]-(b) WHERE id(a) = $nodeId RETURN r"

func GetDirectlyRelatedNodes(nodeId int64, nodeName string) (interface{}, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["nodeId"] = nodeId
	paramsMap["nodeName"] = nodeName
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

var deleteNodeStmtTmpl = `
	MATCH (n:%s)
	WHERE id(n) = $nodeId AND n.name = $nodeName
	DELETE n
`

func DeleteNode(id int64, nodeName, label string) (interface{}, error) {
	deleteNodeStmt := fmt.Sprintf(deleteNodeStmtTmpl, label)
	paramsMap := make(map[string]interface{})
	paramsMap["nodeId"] = id
	paramsMap["nodeName"] = nodeName
	return QueryNeo4j(deleteNodeStmt, paramsMap, false)
}
