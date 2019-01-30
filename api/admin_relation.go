package api

import (
	"errors"
	"strings"

	"github.com/labstack/echo"
	"gitlab.wallstcn.com/matrix/xgbkb/business"
	"gitlab.wallstcn.com/matrix/xgbkb/types"
)

func ApiCreateRelation(ctx echo.Context) (interface{}, error) {
	var relationIn types.RelationIn
	if err := ctx.Bind(&relationIn); err != nil {
		return nil, err
	}

	if relationIn.FromNodeId <= 0 {
		return nil, errors.New("fromNodeId invalid")
	}
	relationIn.FromNodeName = strings.TrimSpace(relationIn.FromNodeName)
	if relationIn.FromNodeName == "" {
		return nil, errors.New("fromNodeName empty")
	}
	if !business.IsValidNodeName(relationIn.FromNodeName) {
		return nil, errors.New("fromNodeName invalid")
	}
	if !relationIn.Relation.IsValid() {
		return nil, errors.New("relation invalid")
	}
	if relationIn.ToNodeId <= 0 {
		return nil, errors.New("toNodeId invalid")
	}
	relationIn.ToNodeName = strings.TrimSpace(relationIn.ToNodeName)
	if relationIn.ToNodeName == "" {
		return nil, errors.New("toNodeName empty")
	}
	if !business.IsValidNodeName(relationIn.ToNodeName) {
		return nil, errors.New("toNodeName invalid")
	}

	return business.CreateRelation(&relationIn)
}
