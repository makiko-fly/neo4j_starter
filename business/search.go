package business

import "fmt"

var searchAllWithNameLikeKeywordStmt = `
	MATCH (n) 
	WHERE n.name =~ $regex 
	RETURN n, labels(n)
	SKIP $offset
	LIMIT $limit
`

func SearchAllWithNameLikeKeywoard(keyword string, page, limit int64) (interface{}, error) {
	statment := ""
	paramsMap := make(map[string]interface{})
	paramsMap["regex"] = fmt.Sprintf(".*%s.*", keyword)
	paramsMap["offset"] = (page - 1) * limit
	paramsMap["limit"] = limit
	return QueryNeo4j(statment, paramsMap, false)
}

var searchInLabelsStmtTmp = `
	MATCH (n)
	WHERE (%s) AND n.name =~ $regex
	RETURN n, labels(n)
	SKIP $offset
	LIMIT $limit
`

func SearchInLabelsWithNameLikeKeyword(keyword string, labels []string, page, limit int64) (interface{}, error) {
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
	paramsMap["offset"] = (page - 1) * limit
	paramsMap["limit"] = limit
	return QueryNeo4j(searchInLabelsStmt, paramsMap, false)
}
