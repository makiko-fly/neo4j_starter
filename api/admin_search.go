package api

import (
	"errors"
	"fmt"
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

	if len(labels) == 0 {
		return business.SearchAllWithNameLikeKeywoard(keywordStr)
	} else {
		return business.SearchInLabelsWithNameLikeKeyword(keywordStr, labels)
	}
}
