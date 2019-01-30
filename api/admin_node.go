package api

import (
	"errors"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"gitlab.wallstcn.com/matrix/xgbkb/business"
)

func ApiGetDirectlyRelatedNodes(ctx echo.Context) (interface{}, error) {
	nodeIdStr := strings.TrimSpace(ctx.QueryParam("nodeId"))
	nodeId, _ := strconv.ParseInt(nodeIdStr, 10, 64)
	if nodeId <= 0 {
		return nil, errors.New("NodeId param is invalid")
	}
	return business.GetDirectlyRelatedNodes(nodeId)
}
