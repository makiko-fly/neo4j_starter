package business

import (
	"fmt"

	"gitlab.wallstcn.com/matrix/xgbkb/common"
	"gitlab.wallstcn.com/matrix/xgbkb/g"
	"gitlab.wallstcn.com/matrix/xgbkb/types"
)

var createRelationStmtTmpl = `
	MATCH (n:%s), (m:%s)
	WHERE id(n) = $fromNodeId and n.name = $fromNodeName and id(m) = $toNodeId and m.name = $toNodeName
	MERGE (n)-[r:%s]->(m)
	RETURN n, r, m
`

func CreateRelation(relationIn *types.RelationIn) (interface{}, error) {
	// infer label from relation
	var fromNodeLabel, toNodeLabel string
	if relationIn.Relation == common.EnumRelation_PRODUCES {
		fromNodeLabel = g.LabelCompany
		toNodeLabel = g.LabelProduct
	} else if relationIn.Relation == common.EnumRelation_HAS_DOWNSTREAM {
		fromNodeLabel = g.LabelProduct
		toNodeLabel = g.LabelProduct
	}
	if fromNodeLabel == "" {
		return nil, fmt.Errorf("Unsupported relation: %v", relationIn.Relation)
	}

	createRelationStmt := fmt.Sprintf(createRelationStmtTmpl, fromNodeLabel, toNodeLabel, relationIn.Relation.String())

	paramsMap := make(map[string]interface{})
	paramsMap["fromNodeId"] = relationIn.FromNodeId
	paramsMap["fromNodeName"] = relationIn.FromNodeName
	paramsMap["toNodeId"] = relationIn.ToNodeId
	paramsMap["toNodeName"] = relationIn.ToNodeName
	return QueryNeo4j(createRelationStmt, paramsMap, true)
}
