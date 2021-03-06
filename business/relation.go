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
	} else if relationIn.Relation == common.EnumRelation_HAS_DOWNSTREAM ||
		relationIn.Relation == common.EnumRelation_HAS_CHILD {
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
	return Neo4jSingleQuery(createRelationStmt, paramsMap, true)
}

var deleteRelationStmtTmpl = `
	MATCH (n:%s)-[r:%s]->(m:%s)
	WHERE id(n) = $fromNodeId AND n.name = $fromNodeName AND id(m) = $toNodeId AND m.name = $toNodeName
	DELETE r
	RETURN r
`

func DeleteRelation(relationIn *types.RelationIn) (interface{}, error) {
	// infer label from relation
	var fromNodeLabel, toNodeLabel string
	if relationIn.Relation == common.EnumRelation_PRODUCES {
		fromNodeLabel = g.LabelCompany
		toNodeLabel = g.LabelProduct
	} else if relationIn.Relation == common.EnumRelation_HAS_DOWNSTREAM ||
		relationIn.Relation == common.EnumRelation_HAS_CHILD {
		fromNodeLabel = g.LabelProduct
		toNodeLabel = g.LabelProduct
	}
	if fromNodeLabel == "" {
		return nil, fmt.Errorf("Unsupported relation: %v", relationIn.Relation)
	}

	deleteRelationStmt := fmt.Sprintf(deleteRelationStmtTmpl, fromNodeLabel, relationIn.Relation.String(), toNodeLabel)

	paramsMap := make(map[string]interface{})
	paramsMap["fromNodeId"] = relationIn.FromNodeId
	paramsMap["fromNodeName"] = relationIn.FromNodeName
	paramsMap["toNodeId"] = relationIn.ToNodeId
	paramsMap["toNodeName"] = relationIn.ToNodeName
	return Neo4jSingleQuery(deleteRelationStmt, paramsMap, false)
}
