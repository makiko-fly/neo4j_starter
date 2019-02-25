package api

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"gitlab.wallstcn.com/matrix/xgbkb/business"
	"gitlab.wallstcn.com/matrix/xgbkb/g"
)

func ApiSearchByName(ctx echo.Context) (interface{}, error) {
	keywordStr := strings.TrimSpace(ctx.QueryParam("keyword"))
	if len(keywordStr) == 0 {
		return nil, errors.New("empty keyword")
	}

	labelsStr := strings.TrimSpace(ctx.QueryParam("labels"))
	labels := strings.Split(labelsStr, ",")
	for _, label := range labels {
		if _, found := g.AllLabelsMap[label]; !found {
			return nil, fmt.Errorf("Label %s invalid", label)
		}
	}

	page, _ := strconv.ParseInt(ctx.Param("page"), 10, 64)
	limit, _ := strconv.ParseInt(ctx.Param("limit"), 10, 64)
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 100
	}
	if limit > 100 {
		return nil, errors.New("Invalid limit")
	}

	if len(labels) == 0 {
		return business.SearchAllWithNameLikeKeywoard(keywordStr, page, limit)
	} else {
		return business.SearchInLabelsWithNameLikeKeyword(keywordStr, labels, page, limit)
	}
}
