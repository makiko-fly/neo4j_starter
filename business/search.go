package business

import "fmt"

// return paging results as well as total
var searchAllWithNameLikeKeywordStmt = `
	MATCH (n) 
	WHERE n.name =~ $regex
	WITH n 
	SKIP $offset LIMIT $limit
	WITH collect(n) as results, collect(labels(n)) as labels
	MATCH (m)
	WHERE m.name =~ $regex 
	WITH results, labels, count(*) as total
	RETURN results, labels, total
`

func SearchAllWithNameLikeKeywoard(keyword string, page, limit int64) (interface{}, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["regex"] = fmt.Sprintf(".*%s.*", keyword)
	paramsMap["offset"] = (page - 1) * limit
	paramsMap["limit"] = limit
	return QueryNeo4j(searchAllWithNameLikeKeywordStmt, paramsMap, false)
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
