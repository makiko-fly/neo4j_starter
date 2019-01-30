package types

import "gitlab.wallstcn.com/matrix/xgbkb/common"

type RelationIn struct {
	FromNodeId   int64               `json:"from_node_id"`
	FromNodeName string              `json:"from_node_name"`
	Relation     common.EnumRelation `json:"relation"`
	ToNodeId     int64               `json:"to_node_id"`
	ToNodeName   string              `json:"to_node_name"`
}
