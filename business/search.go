package business

import "fmt"

func SearchAllWithNameLikeKeywoard(keyword string) (interface{}, error) {
	statment := "MATCH (n) where n.name =~ $regex return n, labels(n)"
	params := map[string]interface{}{
		"regex": fmt.Sprintf(".*%s.*", keyword),
	}
	return QueryNeo4j(statment, params, false)
}

var searchInLabelsStmtTmp = `
	MATCH (n)
	WHERE %s AND n.name =~ $regex
	RETURN n
`

func SearchInLabelsWithNameLikeKeyword(keyword string, labels []string) (interface{}, error) {
	subQuery := fmt.Sprintf("n:%s", labels[0])
	if len(labels) == 1 {
		// pass
	} else {
		for i, label := range labels {
			if i > 0 {
				subQuery += fmt.Sprintf(" OR n:%s", label)
			}
		}
	}

	searchInLabelsStmt := fmt.Sprintf(searchInLabelsStmtTmp, subQuery)

	paramsMap := make(map[string]interface{})
	paramsMap["regex"] = fmt.Sprintf(".*%s.*", keyword)
	return QueryNeo4j(searchInLabelsStmt, paramsMap, false)
}
